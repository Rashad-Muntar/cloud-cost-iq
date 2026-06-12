package ingestion

import "time"

type RawCostRecord struct {
	Provider string // AWS

	CloudAccountID string

	Service string

	Region string

	UsageAmount float64

	CostAmount float64

	Currency string

	UsageStart time.Time

	UsageEnd time.Time
}