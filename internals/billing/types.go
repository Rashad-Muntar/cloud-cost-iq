package billing

import (
	"context"
	"time"

	"github.com/google/uuid"
)


type Repository interface {
	InsertCost(
		ctx context.Context,
		input RecordCostInput,
	) error

	GetDailyCostSummary(
		ctx context.Context,
		date time.Time,
		accountID string,
		service string,
	) (*DailyCostSummary, error)

}

type Service struct {
	repo Repository
}

type AccountFilter struct {
    Environment *string
    IsActive    *bool
    StartDate   *time.Time
    EndDate     *time.Time
}

type CostRecordWithAccount struct {
    CostEvent
    AccountName  string
    Environment  string
    AWSAccountID string
}


type ServiceCost struct {
	Service string  `json:"service"`
	Cost    float64 `json:"cost"`
}

type DailyCostSummary struct {
	TotalCost float64       `json:"total_cost"`
	Services  []ServiceCost `json:"services"`
}

type RecordCostInput struct {
	InternalID uuid.UUID `json:"id"`

	AwsAccountID string `json:"account_id"`

	Service string `json:"service"`
	Region  string `json:"region"`

	CostAmount  float64 `json:"cost_amount"`
	UsageAmount float64 `json:"usage_amount"`

	Currency string `json:"currency"`

	UsageDate time.Time `json:"usage_date"`
	CreatedAt time.Time `json:"created_at"`
	IdempotencyKey string
}



