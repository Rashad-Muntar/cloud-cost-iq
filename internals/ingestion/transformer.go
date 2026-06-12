package ingestion

import (
	"github.com/google/uuid"
)

func Transform(r RawCostRecord, internalAccountID uuid.UUID) NormalizedCostEvent {

	return NormalizedCostEvent{
		AccountID: internalAccountID,

		Service: r.Service,

		Region: r.Region,

		Cost: r.CostAmount,

		Usage: r.UsageAmount,

		Currency: r.Currency,

		Timestamp: r.UsageEnd,
	}
}