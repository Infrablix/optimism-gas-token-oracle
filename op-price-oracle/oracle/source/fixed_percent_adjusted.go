package source

import (
	"context"
	"errors"
	"math/big"
)

var basePercent = big.NewInt(100)

// FixedPercentAdjustedPriceSource is a price source that lowers the price by a certain percentage, effectively
// increasing the fee charged for transactions on L2.
type FixedPercentAdjustedPriceSource struct {
	source                      OraclePriceSource
	adjustmentPercentMultiplier *big.Int
}

func NewFixedPercentAdjustedPriceSource(source OraclePriceSource, adjustmentPercent uint64) (*FixedPercentAdjustedPriceSource, error) {
	if adjustmentPercent >= basePercent.Uint64() {
		return nil, errors.New("adjustment percent must be less than 100")
	}
	return &FixedPercentAdjustedPriceSource{
		source:                      source,
		adjustmentPercentMultiplier: new(big.Int).Sub(basePercent, big.NewInt(int64(adjustmentPercent))),
	}, nil
}

func (s *FixedPercentAdjustedPriceSource) GetPriceInEth(ctx context.Context) (*big.Int, error) {
	price, err := s.source.GetPriceInEth(ctx)
	if err != nil {
		return nil, err
	}

	// if same as base percent, no adjustment needed
	if s.adjustmentPercentMultiplier.Cmp(basePercent) == 0 {
		return price, nil
	}

	// lowers the price by the buffer percent - the lower the price, the higher the fee that's charged
	priceWithBuffer := new(big.Int).Mul(price, s.adjustmentPercentMultiplier)
	priceWithBuffer.Div(priceWithBuffer, basePercent)

	return priceWithBuffer, nil
}
