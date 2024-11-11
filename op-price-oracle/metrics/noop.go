package metrics

import (
	"io"

	"github.com/ethereum-optimism/optimism/op-price-oracle/oracle/bindings"
	txmetrics "github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type noopMetrics struct {
	txmetrics.NoopTxMetrics
}

var NoopMetrics = new(noopMetrics)

func (n *noopMetrics) RecordUp() {}

func (n *noopMetrics) UpdaterStatusIncrement(status string) {}

func (n *noopMetrics) StartBalanceMetrics(l log.Logger, client *ethclient.Client, account common.Address) io.Closer {
	return nil
}

func (n *noopMetrics) StartTokenPriceMetrics(l log.Logger, oracle *bindings.CustomGasTokenPriceOracle) io.Closer {
	return nil
}
