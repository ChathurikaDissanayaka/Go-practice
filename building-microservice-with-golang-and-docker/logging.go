package main

import (
	"context"
	"time"
)

type loggingService struct {
	next PriceFetcher
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {

	}(time.Now())
}

go mod init github.com/ChathurikaDissanayaka/building-microservice-with-golang-and-docker  