package analytics

import "context"

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

func (s *Service) GetCostSummary(
	ctx context.Context,
	query CostQuery,
)(
	*CostSummary,
	error,
){

	return s.repo.GetCostSummary(ctx, query,)
}