package worker

import (
	"context"
	// "fmt"
	"time"

	"github.com/cloud-cost-iq/internals/ingestion"
)

type Processor struct {
	ingestionService *ingestion.Service
	dlq              *DLQ
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
	maxRetries := 3

	var err error

	for i := 0; i < maxRetries; i++ {

		err = p.ingestionService.Ingest(
			ctx,
			job.Payload,
			job.InternalID,
		)

		if err == nil {
			return nil
		}

		time.Sleep(time.Duration(i+1) * time.Second)
	}
	if err != nil {
	p.dlq.Push(job)
	return err
	}

	return err
}