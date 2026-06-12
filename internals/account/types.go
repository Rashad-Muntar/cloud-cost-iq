// internal/account/types.go
package account

import (
    "context"
    "time"
    
    "github.com/google/uuid"
)


type Environment string

const (
	Production  Environment = "production"
	Staging     Environment = "staging"
	Development Environment = "development"
	Security    Environment = "security"
)


type Account struct {
    // Internal ID - for our system use only
    ID uuid.UUID `json:"id"`
    
    // AWS Account ID - immutable, from AWS
    AWSAccountID string `json:"aws_account_id"`
    
    // Human-readable name
    Name string `json:"name"`
    
    // Which environment? dev, staging, prod, sandbox
    Environment Environment `json:"environment"`
    
    // Account status
    IsActive bool `json:"is_active"`
    
    // Metadata
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`  

}


type AccountStatus string

const (
    AccountStatusActive    AccountStatus = "active"    
    AccountStatusDisabled  AccountStatus = "disabled"  
    AccountStatusArchived  AccountStatus = "archived"  
    AccountStatusPending   AccountStatus = "pending"  
)

type Repository interface {
	Create(
		ctx context.Context,
		input Account,
	) (*Account, error)

	GetByAWSAccountID(
		ctx context.Context,
		awsAccountID string,
	) (*Account, error)

	List(
		ctx context.Context,
        filter AccountFilter,
	) ([]*Account, error)

}

type CreateAccountInput struct {
	AWSAccountID string
	Name string
	Environment Environment
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Service struct {
	repo Repository
}

type AccountFilter struct {
    isActive      bool
    Limit       int
    Offset      int
}
