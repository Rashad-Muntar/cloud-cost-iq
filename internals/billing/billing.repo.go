package billing

import (
	"context"
	"time"
	"github.com/cloud-cost-iq/internals/db"
)

type repository struct {
	db *db.Database
}

func NewRepository(database *db.Database) Repository {  // ← returns pointer to private struct
    return &repository{db: database}
}

func (repo *repository) InsertCost(ctx context.Context, input RecordCostInput) error {
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


func (repo *repository) GetDailyCostSummary(
	ctx context.Context,
	date time.Time,
) (*DailyCostSummary, error) {
	query := `
		SELECT service, COALESCE(SUM(cost_amount),0) FROM cost_events
		WHERE DATE(usage_date) = DATE($1)
		GROUP BY service
		ORDER BY SUM(cost_amount) DESC
	`
	rows, _ := repo.db.Pool.Query(ctx, query, date)
	defer rows.Close()

summary := &DailyCostSummary{}
for rows.Next() {

	var service string
	var cost float64

	if err := rows.Scan(
		&service,
		&cost,
	); err != nil {
		return nil, err
	}

	summary.Services = append(
		summary.Services,
		ServiceCost{
			Service: service,
			Cost: cost,
		},
	)

	summary.TotalCost += cost
}
return summary, nil
}