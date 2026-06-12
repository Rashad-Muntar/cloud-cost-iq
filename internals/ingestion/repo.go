package ingestion

import (
	"context"

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
		billing.CostRecord{
			AccountID: event.AccountID,
			Service: event.Service,
			Region: event.Region,
			CostAmount: event.Cost,
			UsageAmount: event.Usage,
			Currency: event.Currency,
			UsageDate: event.Timestamp,
		},
	)
}