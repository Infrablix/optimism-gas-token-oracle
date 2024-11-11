package oracle

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum-optimism/optimism/op-price-oracle/metrics"
	"github.com/ethereum-optimism/optimism/op-price-oracle/oracle/bindings"
	"github.com/ethereum-optimism/optimism/op-price-oracle/oracle/impls"
	"github.com/ethereum-optimism/optimism/op-price-oracle/oracle/source"
	"github.com/ethereum-optimism/optimism/op-service/httputil"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/urfave/cli/v2"

	"github.com/ethereum-optimism/optimism/op-price-oracle/flags"
	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
)

// Main is the entrypoint into the Batch Submitter.
// This method returns a cliapp.LifecycleAction, to create an op-service CLI-lifecycle-managed batch-submitter with.
func Main() cliapp.LifecycleAction {
	return func(cliCtx *cli.Context, closeApp context.CancelCauseFunc) (cliapp.Lifecycle, error) {
		cfg := NewConfig(cliCtx)
		if err := cfg.Check(); err != nil {
			return nil, fmt.Errorf("invalid CLI flags: %w", err)
		}

		logger := oplog.NewLogger(oplog.AppOut(cliCtx), cfg.LogConfig)
		oplog.SetGlobalLogHandler(logger.Handler())
		opservice.ValidateEnvVars(flags.EnvVarPrefix, flags.Flags, logger)

		logger.Info("Initializing Price Oracle")
		return oracleServiceFromCLIConfig(cliCtx.Context, cfg, logger)
	}
}

type priceOracleService struct {
	config   *CLIConfig
	l2Client *ethclient.Client
	oracle   *bindings.CustomGasTokenPriceOracle
	source   source.OraclePriceSource
	txmgr    txmgr.TxManager
	logger   log.Logger
	metrics  metrics.Metricer

	ctx      context.Context
	stopping atomic.Bool
	stopped  atomic.Bool

	metricsSrv         *httputil.HTTPServer
	balanceMetricer    io.Closer
	tokenPriceMetricer io.Closer
}

func (s *priceOracleService) Start(_ context.Context) error {
	if s.config.MetricsConfig.Enabled {
		s.logger.Debug("Starting metrics server", "addr", s.config.MetricsConfig.ListenAddr, "port", s.config.MetricsConfig.ListenPort)
		m, ok := s.metrics.(opmetrics.RegistryMetricer)
		if !ok {
			return fmt.Errorf("metrics were enabled, but metricer %T does not expose registry for metrics-server", s.metrics)
		}

		metricsSrv, err := opmetrics.StartServer(m.Registry(), s.config.MetricsConfig.ListenAddr, s.config.MetricsConfig.ListenPort)
		if err != nil {
			return fmt.Errorf("failed to start metrics server: %w", err)
		}
		s.logger.Info("Started metrics server", "addr", metricsSrv.Addr())

		s.metricsSrv = metricsSrv
		s.metrics.RecordUp()
		s.metrics.StartBalanceMetrics(s.logger, s.l2Client, s.txmgr.From())
		s.metrics.StartTokenPriceMetrics(s.logger, s.oracle)
	}

	go s.loop()

	return nil
}

func (s *priceOracleService) Stop(_ context.Context) error {
	s.stopping.Store(true)

	var result error
	if s.balanceMetricer != nil {
		if err := s.balanceMetricer.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close balance metricer: %w", err))
		}
	}

	if s.tokenPriceMetricer != nil {
		if err := s.tokenPriceMetricer.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close token price metricer: %w", err))
		}
	}

	if s.metricsSrv != nil {
		if err := s.metricsSrv.Stop(context.Background()); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to stop metrics server: %w", err))
		}
	}

	if s.l2Client != nil {
		s.l2Client.Close()
	}

	return result
}

func (s *priceOracleService) Stopped() bool {
	return s.stopped.Load()
}

func oracleServiceFromCLIConfig(ctx context.Context, config *CLIConfig, logger log.Logger) (*priceOracleService, error) {
	l2Client, err := ethclient.DialContext(ctx, config.L2EthRpc)
	if err != nil {
		return nil, err
	}

	oracle, err := bindings.NewCustomGasTokenPriceOracle(predeploys.CustomGasTokenPriceOracleAddr, l2Client)
	if err != nil {
		return nil, err
	}

	var oracleMetrics metrics.Metricer = metrics.NoopMetrics
	if config.MetricsConfig.Enabled {
		oracleMetrics = metrics.NewMetrics("")
	}

	// IMPORTANT: set the L1 RPC URL to the L2 RPC URL, so the tx manager will submit to L2
	config.TxMgrConfig.L1RPCURL = config.L2EthRpc
	mgr, err := txmgr.NewSimpleTxManager("custom-gas-token-price-oracle", logger, oracleMetrics, config.TxMgrConfig)
	if err != nil {
		return nil, err
	}

	chainId, err := l2Client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	service := &priceOracleService{
		config:   config,
		l2Client: l2Client,
		oracle:   oracle,
		source:   impls.NewOracleFromChainId(chainId.Uint64(), logger, config.L1EthRpc),
		txmgr:    mgr,
		logger:   logger,
		metrics:  oracleMetrics,
		ctx:      ctx,
		stopping: atomic.Bool{},
		stopped:  atomic.Bool{},
	}

	return service, nil
}

func (s *priceOracleService) loop() {
	status, err := s.oracle.GetPrice(nil)
	if err != nil {
		s.logger.Crit("failed to get initial oracle status", "err", err)
	}

	minPriceChangePercent := big.NewFloat(s.config.MinPriceChangePercent)
	lastPrice := status.PriceInEth
	lastUpdateTimestamp := status.LastUpdateTimestamp.Uint64()
	for !s.stopping.Load() {
		time.Sleep(2 * time.Second)

		now := uint64(time.Now().Unix())
		price, err := s.source.GetPriceInEth(s.ctx)

		if err != nil {
			s.metrics.UpdaterStatusIncrement("source-call-error")
			s.logger.Error("failed to get new oracle price", "err", err)
			continue
		}

		if price.Cmp(common.Big0) == 0 {
			s.metrics.UpdaterStatusIncrement("source-price-zero")
			s.logger.Warn("new oracle price is zero", "last", lastPrice.String())
			continue
		}

		shouldUpdate := lastPrice.Cmp(common.Big0) == 0
		if !shouldUpdate {
			lastPriceFloat := new(big.Float).SetInt(lastPrice)
			priceFloat := new(big.Float).SetInt(price)
			priceChange := new(big.Float).Sub(priceFloat, lastPriceFloat)
			priceChange = priceChange.Abs(priceChange)
			priceChangePercent := new(big.Float).Quo(priceChange, lastPriceFloat)
			priceChangePercent = priceChangePercent.Mul(priceChangePercent, new(big.Float).SetFloat64(100))
			shouldUpdate = priceChangePercent.Cmp(minPriceChangePercent) > 0
			s.logger.Info("new oracle price", "new", price.String(), "last", lastPrice.String(), "diff_percent", priceChangePercent.String())
		}

		if !shouldUpdate && s.config.MaxElapsedTime > 0 {
			shouldUpdate = (now - lastUpdateTimestamp) > uint64(s.config.MaxElapsedTime.Seconds())
		}

		if !shouldUpdate {
			continue
		}

		txData := []byte{0x82, 0xab, 0x89, 0x0a}
		txData = append(txData, common.BigToHash(price).Bytes()...)

		receipt, err := s.txmgr.Send(s.ctx, txmgr.TxCandidate{
			TxData:   txData,
			Blobs:    nil,
			To:       &predeploys.CustomGasTokenPriceOracleAddr,
			GasLimit: 125_000,
			Value:    nil,
		})

		if err != nil {
			s.metrics.UpdaterStatusIncrement("update-tx-send-error")
			s.logger.Error("failed to send oracle update transaction", "err", err)
			continue
		}

		if receipt.Status != 1 {
			s.metrics.UpdaterStatusIncrement("update-tx-revert")
			s.logger.Error("oracle update transaction FAIL", "tx", receipt.TxHash.String())
			continue
		}

		s.metrics.UpdaterStatusIncrement("update-success")
		s.logger.Info("oracle update transaction SUCCESS", "tx", receipt.TxHash.String())

		lastPrice = price
		lastUpdateTimestamp = now
	}

	s.logger.Info("Price Oracle stopped")
	s.stopped.Store(true)
}
