package ingestion

import (
	"context"
	// "fmt"

	"github.com/cloud-cost-iq/internals/billing"
)

type repository struct {
	billingRepo billing.Repository
}

func NewRepository(b billing.Repository) Repository {
	return &repository{
		billingRepo: b,
	}
}

func (r *repository) StoreCostEvent(
	ctx context.Context,
	event NormalizedCostEvent,
) error {
	return r.billingRepo.InsertCost(
		ctx,
		billing.RecordCostInput{
			InternalID: event.InternalID,
			AwsAccountID: event.AwsAccountID,
			Service: event.Service,
			Region: event.Region,
			CostAmount: event.Cost,
			UsageAmount: event.Usage,
			Currency: event.Currency,
			UsageDate: event.Timestamp,
			IdempotencyKey: event.IdempotencyKey,
		},
	)
}