package billing

import (
	"time"
	"github.com/google/uuid"
)

type CostRecord struct {
	InternalID          uuid.UUID `db:"id"`
    AwsAccountID   uuid.UUID `db:"account_id"` 
    Service     string    `db:"service"`
    Region      string    `db:"region"`
    CostAmount  float64   `db:"cost_amount"`
    UsageAmount float64   `db:"usage_amount"`
    Currency    string    `db:"currency"`
    UsageDate   time.Time `db:"usage_date"`
    CreatedAt   time.Time `db:"created_at"`
    IdempotencyKey string `db:"idempotency_key"`
}