package impls

import (
	"github.com/ethereum-optimism/optimism/op-price-oracle/oracle/source"
	"github.com/ethereum/go-ethereum/log"
)

func NewOracleFromChainId(chainId uint64, logger log.Logger, l1RpcUrl string) source.OraclePriceSource {
	switch chainId {
	default:
		logger.Crit("unsupported chainId", "chainId", chainId)
	}
	return nil
}
