package source

import (
	"context"
	"math/big"
)

type DerivedPriceSource struct {
	sourceA OraclePriceSource // TKN / CUSTOM_GAS_TOKEN
	sourceB OraclePriceSource // WETH / TKN
	scalar  *big.Int          // TKN decimals
}

func NewDerivedPriceSource(sourceA, sourceB OraclePriceSource, scalar *big.Int) *DerivedPriceSource {
	return &DerivedPriceSource{
		sourceA: sourceA,
		sourceB: sourceB,
		scalar:  scalar,
	}
}

func (s *DerivedPriceSource) GetPriceInEth(ctx context.Context) (*big.Int, error) {
	priceA, err := s.sourceA.GetPriceInEth(ctx)
	if err != nil {
		return nil, err
	}

	priceB, err := s.sourceB.GetPriceInEth(ctx)
	if err != nil {
		return nil, err
	}

	price := new(big.Int).Mul(priceA, priceB)
	price.Div(price, s.scalar)

	return price, nil
}
