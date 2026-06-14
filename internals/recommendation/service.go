package recommendation

import "context"

type Service struct {
	engine *Engine

	repo Repository
}

func NewService(
	engine *Engine,

	repo Repository,
) *Service {

	return &Service{
		engine: engine,
		repo: repo,
	}
}


func (s *Service) Generate(
	ctx context.Context,
	accountID string,
) error {

	recs,
	err :=
	s.engine.
	Generate(
		ctx,
		accountID,
	)

	if err != nil {
		return err
	}

	for _,
	r :=
	range recs {

		err =
		s.repo.
		Create(
			ctx,
			r,
		)

		if err != nil {
			return err
		}
	}

	return nil
}