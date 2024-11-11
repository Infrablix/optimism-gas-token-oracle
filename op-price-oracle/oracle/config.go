package oracle

import (
	"errors"
	"time"

	"github.com/ethereum-optimism/optimism/op-price-oracle/flags"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	opmetrics "github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/urfave/cli/v2"
)

type CLIConfig struct {
	// L1EthRpc is the URL of the L1 execution client RPC server.
	L1EthRpc string
	// L2EthRpc is the URL of the L2 execution client RPC server.
	L2EthRpc string

	// MinPriceChangePercent is the minimum percentage change in price that will trigger a new price update. Value of
	// 0 means that every price update will trigger a new oracle update.
	MinPriceChangePercent float64

	// MaxElapsedTime is the maximum amount of time that can pass before a new price update is triggered, even if
	// MinPriceChangePercent is not reached. Value of 0 disables this feature.
	MaxElapsedTime time.Duration

	TxMgrConfig   txmgr.CLIConfig
	MetricsConfig opmetrics.CLIConfig
	LogConfig     oplog.CLIConfig
}

func (cli CLIConfig) Check() error {
	if cli.L1EthRpc == "" {
		return errors.New("L1EthRpc must be set")
	}
	if cli.L2EthRpc == "" {
		return errors.New("L2EthRpc must be set")
	}
	if cli.MinPriceChangePercent < 0 {
		return errors.New("MinPriceChangePercent must be non-negative")
	}
	if cli.MaxElapsedTime < 0 {
		return errors.New("MaxElapsedTime must be positive or zero (disables this feature)")
	}
	if err := cli.TxMgrConfig.Check(); err != nil {
		return err
	}
	if err := cli.MetricsConfig.Check(); err != nil {
		return err
	}
	return nil
}

func NewConfig(ctx *cli.Context) *CLIConfig {
	return &CLIConfig{
		L1EthRpc:              ctx.String(flags.L1EthRpcFlag.Name),
		L2EthRpc:              ctx.String(flags.L2EthRpcFlag.Name),
		MinPriceChangePercent: ctx.Float64(flags.MinPriceChangeFlag.Name),
		MaxElapsedTime:        ctx.Duration(flags.MaxElapsedTimeFlag.Name),
		MetricsConfig:         opmetrics.ReadCLIConfig(ctx),
		TxMgrConfig:           txmgr.ReadCLIConfig(ctx),
		LogConfig:             oplog.ReadCLIConfig(ctx),
	}
}
