package account

import "time"

type AccountModel struct {
	ID string

	AWSAccountID string

	Name string

	Environment string

	IsActive bool

	CreatedAt time.Time
	UpdatedAt time.Time
}