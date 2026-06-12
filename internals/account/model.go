package account

import (
	"time" 
"github.com/google/uuid"
)

type AccountModel struct {
    ID           uuid.UUID  `db:"id"`
    AWSAccountID string     `db:"aws_account_id"`
    Name         string     `db:"name"`
    Environment  string     `db:"environment"` // development, staging, production
    IsActive     bool       `db:"is_active"`
    DeletedAt    *time.Time `db:"deleted_at"`
    CreatedAt    time.Time  `db:"created_at"`
    UpdatedAt    time.Time  `db:"updated_at"`

}