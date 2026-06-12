package ingestion

import (
	"context"
	"time"
	"github.com/google/uuid"
)


type NormalizedCostEvent struct {
	AccountID uuid.UUID

	Service string

	Region string

	Cost float64

	Usage float64

	Currency string

	Timestamp time.Time
}

type Repository interface {
	StoreCostEvent(
		ctx context.Context,
		event NormalizedCostEvent,
	) error
}