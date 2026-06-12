// internal/account/service.go
package account

import (
	"context"
	"fmt"
	"time"
    "regexp"
	"github.com/google/uuid"
)


func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, input CreateAccountInput) (*Account, error) {
    // Validation - senior reasoning: fail fast with clear errors
    if input.AWSAccountID == "" {
        return nil, fmt.Errorf("aws_account_id is required")
    }

    matched, _ := regexp.MatchString(
	`^\d{12}$`,
	input.AWSAccountID,
)

if !matched {
	return nil,  fmt.Errorf("aws_account_id is invalid")
}
    
    if input.Name == "" {
        return nil, fmt.Errorf("account name is required")
    }
    
    // Check if account already exists
    existing, _ := s.repo.GetByAWSAccountID(ctx, input.AWSAccountID)
    if existing != nil {
        return nil, fmt.Errorf("account with AWS ID %s already exists", input.AWSAccountID)
    }
    
    now := time.Now()
    
    account := Account{
        ID:          uuid.New(),
        AWSAccountID:        input.AWSAccountID,
        Name:                input.Name,
        Environment:         input.Environment,
        CreatedAt:           now,
        UpdatedAt:           now,
    }
    

    
   created, err := s.repo.Create(ctx, account)
    if err != nil {
        return nil, fmt.Errorf("failed to create account: %w", err)
    }
    
    return created, nil
}



func (s *Service) GetAccountByAWSID(ctx context.Context, awsID string) (*Account, error) {
    return s.repo.GetByAWSAccountID(ctx, awsID)
}

func (s *Service) List(ctx context.Context, filter AccountFilter) ([]*Account, error) {
    // Set reasonable defaults
    if filter.Limit == 0 || filter.Limit > 100 {
        filter.Limit = 50
    }
    
    return s.repo.List(ctx, filter)
}
