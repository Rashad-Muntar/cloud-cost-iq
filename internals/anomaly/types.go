package anomaly

import (
	"context"
)

type Service struct {
	repo Repository
	analytics Analytics
}


type Analytics interface {
	HistoricalDailyCosts(
		ctx context.Context,
		accountID string,
		service string,
	)([]float64, error,)

	TodayCost(
		ctx context.Context,
		accountID string,
		service string,
	)(float64,error,
	)
}