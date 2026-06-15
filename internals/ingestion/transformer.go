package ingestion

import "github.com/google/uuid"

func Transform(r RawCostEvent, internalAccountID uuid.UUID) NormalizedCostEvent {

	return NormalizedCostEvent{
		InternalID: internalAccountID,
		AwsAccountID: r.AwsAccountID,
		Service: r.Service,
		Region: r.Region,
		Cost: r.CostAmount,
		Usage: r.UsageAmount,
		Currency: r.Currency,
		Timestamp: r.UsageEnd,
		IdempotencyKey: GenerateIdempotencyKey(r),
	}
}