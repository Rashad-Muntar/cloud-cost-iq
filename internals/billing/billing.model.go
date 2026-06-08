package billing

import "time"

type CostRecord struct {
	ID string

	AccountID string

	Service string
	Region  string

	CostAmount  float64
	UsageAmount float64

	Currency string

	UsageDate time.Time
	CreatedAt time.Time
}