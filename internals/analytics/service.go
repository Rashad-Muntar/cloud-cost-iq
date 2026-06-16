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

func (s *Service,) HistoricalDailyCosts(
	ctx context.Context,
	accountID string,
	service string,
)([]float64, error,){

	return s.repo.HistoricalDailyCosts(ctx, accountID,service,)
}

func (s *Service,) TodayCost(
	ctx context.Context,

	accountID string,

	service string,
)(float64, error,
){

	return s.repo.TodayCost(ctx,accountID,service,)
}