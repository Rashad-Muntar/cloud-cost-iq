// internal/account/repository.go
package account

import (
    "context"
    "fmt"
    "github.com/cloud-cost-iq/internals/db"

)

type repository struct {
    db *db.Database
}

func NewRepository(database *db.Database) Repository {
    return &repository{db: database}
}

func (r *repository) Create(ctx context.Context, input Account) (*Account, error) {
    query := `
        INSERT INTO accounts (
            id, name, aws_account_id, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5)
    `
    var account Account
    _, err := r.db.Pool.Exec(ctx, query,
        input.InternalID,
        input.Name,
        input.AwsAccountID,
        input.CreatedAt,
        input.UpdatedAt,

    )
    
    return &account, err
}

func (r *repository) GetByAWSAccountID(ctx context.Context, awsID string) (*Account, error) {
    query := `
        SELECT 
            id, aws_account_id, name, environment, is_active, created_at, updated_at,
        FROM accounts
        WHERE aws_account_id = $1
    `
    var account Account
    err := r.db.Pool.QueryRow(ctx, query, awsID).Scan(
        &account.InternalID,
        &account.AwsAccountID,
        &account.Name,
        &account.Environment,
        &account.IsActive,
        &account.CreatedAt,
        &account.UpdatedAt,
    )
    
    if err != nil {
        return nil, err
    }
    
    return &account, nil
}

func (r *repository) List(ctx context.Context, filter AccountFilter) ([]*Account, error) {
    query := `
        SELECT 
            internal_id, aws_account_id, name,
            environment, owner_team, is_active, , created_at, updated_at,
        FROM accounts
    `
    args := []interface{}{}
    argPosition := 1

    query += " ORDER BY name ASC"
    
    if filter.Limit > 0 {
        query += fmt.Sprintf(" LIMIT $%d", argPosition)
        args = append(args, filter.Limit)
        argPosition++
        
        if filter.Offset > 0 {
            query += fmt.Sprintf(" OFFSET $%d", argPosition)
            args = append(args, filter.Offset)
        }
    }
    
    rows, err := r.db.Pool.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var accounts []*Account
    for rows.Next() {
        var account Account
        err := rows.Scan(
            &account.InternalID,
            &account.AwsAccountID,
            &account.Name,
            &account.Environment,
            &account.IsActive,
            &account.CreatedAt,
            &account.UpdatedAt,

        )
        if err != nil {
            return nil, err
        }

        accounts = append(accounts, &account)
    }
    
    return accounts, nil
}






