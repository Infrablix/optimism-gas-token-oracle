package source

import (
	"context"
	"math/big"
)

type OraclePriceSource interface {
	// GetPriceInEth returns the price of the token in ETH.
	GetPriceInEth(ctx context.Context) (*big.Int, error)
}
