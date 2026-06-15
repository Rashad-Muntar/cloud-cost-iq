package recommendation

import (
	"context"

	"github.com/google/uuid"

	"github.com/cloud-cost-iq/internals/db"
)

type Repository interface {

	Create(
		ctx context.Context,
		r Recommendation,
	) error

	List(
		ctx context.Context,
		accountID string,
	)(
		[]Recommendation,
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

func (r *repository) Create(
	ctx context.Context,
	rec Recommendation,
) error {

	query := `
	INSERT INTO recommendations
	(
		id,
		account_id,
		type,
		title,
		description,
		estimated_monthly_savings,
		severity,
		status
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6,$7,$8
	)
	`

	_, err :=
		r.db.Pool.Exec(
			ctx,
			query,

			uuid.New().String(),

			rec.AccountID,

			rec.Type,

			rec.Title,

			rec.Description,

			rec.EstimatedMonthlySavings,

			rec.Severity,

			rec.Status,
		)

	return err
}

func (r *repository) List(
	ctx context.Context,
	accountID string,
)(
	[]Recommendation,
	error,
){

	query := `
	SELECT
		id,
		account_id,
		type,
		title,
		description,
		estimated_monthly_savings,
		severity,
		status,
		created_at
	FROM recommendations
	WHERE account_id=$1
	ORDER BY created_at DESC
	`
	rows, err := r.db.Pool.Query(ctx, query, accountID)

	if err != nil {
		return nil,
		err
	}

	defer rows.Close()
	var output []Recommendation

	for rows.Next(){
		var x Recommendation
		rows.Scan(
			&x.ID,
			&x.AccountID,
			&x.Type,
			&x.Title,
			&x.Description,
			&x.EstimatedMonthlySavings,
			&x.Severity,
			&x.Status,
			&x.CreatedAt,
		)

		output = append(output, x)
	}

	return output,nil
}