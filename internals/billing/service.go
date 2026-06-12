package billing

import (
	"context"
	"time"
	"strings"
	"github.com/google/uuid"
)


func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) RecordCost(ctx context.Context, accountID uuid.UUID, service, region string, cost, usage float64,) error {
	normalizedService := strings.ToLower(service)
	record := CostRecord{
		ID: uuid.New().String(),

		AccountID: accountID,
		Service: normalizedService,
		Region: region,

		CostAmount:  cost,
		UsageAmount: usage,

		Currency: "USD",

		UsageDate: time.Now(),
		CreatedAt: time.Now(),
	}
	
	return s.repo.InsertCost(ctx, record)
}

func (s *Service) GetDailySummary(ctx context.Context, date time.Time, accountID, service string) (*DailyCostSummary, error) {
	normalizedService := strings.ToLower(service)
	normalizedAccountID := strings.ToLower(accountID)
	return s.repo.GetDailyCostSummary(ctx, date, normalizedAccountID, normalizedService)
}