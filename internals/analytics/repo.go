package analytics

import (
	"context"

	"github.com/cloud-cost-iq/internals/db"
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

func (r *repository) GetCostSummary(
	ctx context.Context,
	query CostQuery,
) (
	*CostSummary,
	error,
) {

	sql := `
	SELECT
		service,
		COALESCE(
			SUM(cost_amount),
			0
		)
	FROM cost_events
	WHERE account_id=$1
	AND usage_date
	BETWEEN $2 AND $3
	GROUP BY service
	ORDER BY SUM(cost_amount)
	DESC
	`

	rows, err :=
		r.db.Pool.Query(
			ctx,
			sql,
			query.AwsAccountID,
			query.From,
			query.To,
		)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result :=
		&CostSummary{}

	for rows.Next() {

		var s ServiceBreakdown

		err :=
			rows.Scan(
				&s.Service,
				&s.Cost,
			)

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

	return result, nil
}