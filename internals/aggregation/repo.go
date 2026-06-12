package aggregation

import (
	"context"
	"log"
	"time"

	"github.com/cloud-cost-iq/internals/db"
)


type repository struct {
	db *db.Database
}

func NewRepository(database *db.Database) Repository {  // ← returns pointer to private struct
    return &repository{db: database}
}

type Repository interface {
	BuildDailySummaries(
		ctx context.Context,
		date time.Time,
	) error
}

func (repo *repository) BuildDailySummaries(
	ctx context.Context,
	date time.Time,
) error {
	query := `
        INSERT INTO daily_cost_summaries (
            summary_date,
            service,
            total_cost
        )
			  SELECT
            DATE(usage_date),
            service,
            SUM(cost_amount)
        FROM cost_events
        WHERE DATE(usage_date) = DATE($1)
        GROUP BY
            DATE(usage_date),
            service
        ON CONFLICT (
            summary_date,
            service
        )
			    DO UPDATE
        SET total_cost = EXCLUDED.total_cost
    `
	result, err := repo.db.Pool.Exec(ctx, query, date)
    if err != nil {
        log.Fatal(err)
    }

    rowsAffected := result.RowsAffected()
    log.Printf("Rows affected: %d", rowsAffected)
	return nil
}