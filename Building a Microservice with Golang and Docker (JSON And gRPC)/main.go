package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, tikcer string)(float64, error){
	return MockPriceFetcher(ctx, tikcer)
}

var priceMocks = map[string]float64{
	"BTC": 20_000_0,
	"ETH": 200,
	"GG": 100_000.0,
}

func MockPriceFetcher(ctx context.Context, tikcer string) (float64, error){
	price, ok := priceMocks[tikcer]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", tikcer)
	}
	return price, nil
}