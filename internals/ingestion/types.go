package ingestion

import (
	"context"
	"time"

	"github.com/google/uuid"
)


type NormalizedCostEvent struct {
	InternalID uuid.UUID
	AwsAccountID string
	Service string
	Region string
	Cost float64
	Usage float64
	Currency string
	Timestamp time.Time
	IdempotencyKey string
}

type Repository interface {
	StoreCostEvent(
		ctx context.Context,
		event NormalizedCostEvent,
	) error
}