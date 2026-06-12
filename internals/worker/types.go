package worker

import (
	"github.com/cloud-cost-iq/internals/ingestion"
	"github.com/google/uuid"
)

type Job struct {
	InternalID uuid.UUID
	Payload ingestion.RawCostRecord
}