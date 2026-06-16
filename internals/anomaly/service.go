package anomaly

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloud-cost-iq/internals/analytics"
)

func NewService(repo Repository, analytics analytics.Repository) *Service {
	return &Service{repo: repo,  analytics: analytics,}
}

func (s *Service) Run(
	ctx context.Context,
	accountID string,
	service string,
) error {
	
	normalizeTService := strings.ToUpper(service)
	fmt.Println("FROM Anomalry", accountID,service)
	history, err := s.analytics.HistoricalDailyCosts(ctx, accountID, normalizeTService) 
	// print("History", history)
	if err != nil {
		fmt.Println("THERES IS ERROR")
		return err
	}

	current, err :=s.analytics.TodayCost(ctx, accountID, normalizeTService,)
	print("Current", history)
	if err != nil {
		return err
	}

	a, ok := Detect(accountID, normalizeTService, history, current,)

	if !ok {

		return nil
	}

	return s.repo.Create(ctx, *a,)
}