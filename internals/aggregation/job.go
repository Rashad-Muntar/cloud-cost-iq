package aggregation

import (
	"context"
	"time"
)

type Job struct {
	service *Service
}

func (j *Job) Execute(
	ctx context.Context,
) error {

	return j.service.
		BuildDailySummaries(
			ctx,
			time.Now(),
		)
}