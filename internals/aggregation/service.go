package aggregation

import (
	"context"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(
	repo Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) BuildDailySummaries(ctx context.Context, time time.Time) error {
	return s.repo.BuildDailySummaries(ctx, time)
}