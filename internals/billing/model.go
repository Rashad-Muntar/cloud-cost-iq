package billing

import (
	"time"
	"github.com/google/uuid"
)

type CostEvent struct {
	InternalID          uuid.UUID `db:"id"`
    AwsAccountID   uuid.UUID `db:"account_id"` 
    Service     string    `db:"service"`
    Region      string    `db:"region"`
    CostAmount  float64   `db:"total_cost"`
    UsageAmount float64   `db:"total_usage"`
    Currency    string    `db:"currency"`
    UsageDate   time.Time `db:"usage_date"`
    CreatedAt   time.Time `db:"created_at"`
    IdempotencyKey string `db:"idempotency_key"`
}

