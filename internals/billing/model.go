package billing

import (
	"time"
	"github.com/google/uuid"
)

type CostRecord struct {
	ID string

	AccountID   uuid.UUID

	Service string
	Region  string

	CostAmount  float64
	UsageAmount float64

	Currency string

	UsageDate time.Time
	CreatedAt time.Time
}