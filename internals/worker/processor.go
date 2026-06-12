package worker

import (
	"context"
	"github.com/cloud-cost-iq/internals/ingestion"
)

type Processor struct {
	ingestionService *ingestion.Service
}

func NewProcessor(
	service *ingestion.Service,
) *Processor {

	return &Processor{
		ingestionService: service,
	}
}

func (p *Processor) Process(
	ctx context.Context,
	job Job,
) error {

	return p.ingestionService.Ingest(
		ctx,
		job.Payload,
		job.InternalID,
	)
}