package ingestion

import (
	"context"
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Ingest(
	ctx context.Context,
	raw RawCostRecord,
	internalAccountID uuid.UUID,
) error {

	event := Transform(raw, internalAccountID)

	return s.repo.StoreCostEvent(ctx, event)
}