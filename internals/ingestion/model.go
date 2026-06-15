package ingestion

import (
	"time"

)

type RawCostEvent struct {
	Provider string // AWS

	AwsAccountID string

	Service string

	Region string

	UsageAmount float64

	CostAmount float64

	Currency string

	UsageStart time.Time

	UsageEnd time.Time
	IdempotencyKey string
}