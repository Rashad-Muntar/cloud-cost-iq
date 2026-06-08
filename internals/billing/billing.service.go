package billing

import (
	"context"
	"time"

	"github.com/google/uuid"
)


func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateCost(
	ctx context.Context,
	accountID, service, region string,
	cost, usage float64,
) error {

	record := CostRecord{
		ID: uuid.New().String(),

		AccountID: accountID,
		Service: service,
		Region: region,

		CostAmount:  cost,
		UsageAmount: usage,

		Currency: "USD",

		UsageDate: time.Now(),
		CreatedAt: time.Now(),
	}

	return s.repo.InsertCost(ctx, record)
}