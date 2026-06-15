package analytics

import (
	"context"
	"fmt"

	"github.com/cloud-cost-iq/internals/db"
	"github.com/google/uuid"
)

type Repository interface {
	GetCostSummary(
		ctx context.Context,
		query CostQuery,
	) (
		*CostSummary,
		error,
	)
}

type repository struct {
	db *db.Database
}

func NewRepository(
	db *db.Database,
) Repository {

	return &repository{
		db: db,
	}
}

func (r *repository) GetCostSummary(ctx context.Context, query CostQuery,) (
	*CostSummary,
	error,
) {
	fmt.Println("Query", query)
	var internalAccountID uuid.UUID
	accountQuery := `SELECT id FROM accounts WHERE aws_account_id = $1`
	err := r.db.Pool.QueryRow(ctx, accountQuery, query.AwsAccountID).Scan(&internalAccountID)
	if err != nil {
		return nil, fmt.Errorf("account not found for aws_account_id %s: %w", query.AwsAccountID, err)
	}
	fmt.Println(internalAccountID)

		sql := `
		SELECT service, COALESCE(SUM(cost_amount), 0)
		FROM cost_events
		WHERE account_id = $1
		AND usage_date BETWEEN $2 AND $3
		GROUP BY service
		ORDER BY SUM(cost_amount) DESC
	`

	rows, err := r.db.Pool.Query(
			ctx,
			sql,
			internalAccountID,
			query.From,
			query.To,
		)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := &CostSummary{}
	
	for rows.Next() {

		var s ServiceBreakdown
		err := rows.Scan(&s.Service, &s.Cost)

		if err != nil {
			return nil, err
		}

		result.Breakdown =
			append(
				result.Breakdown,
				s,
			)

		result.TotalCost += s.Cost
	}

	result.InternalID = internalAccountID
	return result, nil
}