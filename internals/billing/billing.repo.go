package billing

import (
	"context"
	"github.com/cloud-cost-iq/internals/db"
)


func NewRepository(database *db.Database) *Repository {
	return &Repository{
		db: database,
	}
}

func (repo *Repository) InsertCost(ctx context.Context, input CostRecord) error {
	query := `
		INSERT INTO cost_events (
			id, account_id, service, region,
			cost_amount, usage_amount,
			currency, usage_date, created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`
	_, err := repo.db.Pool.Exec(ctx, query,
		input.ID,
		input.AccountID,
		input.Service,
		input.Region,
		input.CostAmount,
		input.UsageAmount,
		input.Currency,
		input.UsageDate,
		input.CreatedAt,
	)

	return err
}