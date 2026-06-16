package anomaly

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/cloud-cost-iq/internals/db"
)

type Repository interface {
	Create(ctx context.Context, a Anomaly,) error

	GetAnomalyByAccount(ctx context.Context, accountID string,)([]Anomaly, error,)
}

type repository struct {
	db *db.Database
}

func NewRepository(
	db *db.Database,
) Repository {
	return &repository{db: db,}
}


func (r *repository) Create(
	ctx context.Context,
	anomaly Anomaly,
) error {
	fmt.Println(anomaly)
	var internalAccountID uuid.UUID
	accountQuery := `SELECT id FROM accounts WHERE aws_account_id = $1`
	err := r.db.Pool.QueryRow(ctx, accountQuery, anomaly.AccountID).Scan(&internalAccountID)
	if err != nil {
		return nil
	}
	fmt.Println(internalAccountID)
	query := `INSERT INTO anomalies(id, account_id, service, expected_cost, actual_cost, deviation, severity)
			VALUES($1,$2,$3,$4,$5,$6,$7)
			`

	_, err = r.db.Pool.Exec(ctx,query, uuid.New().String(),
		internalAccountID,
		anomaly.Service,
		anomaly.ExpectedCost,
		anomaly.ActualCost,
		anomaly.Deviation,
		anomaly.Severity,
	)

	return err
}

func(r *repository) GetAnomalyByAccount(ctx context.Context, accountID string,) ([]Anomaly, error){
	query := `SELECT service, expected_cost, actual_cost, deviation, severity, detected_at FROM anomalies WHERE account_id=$1 FROM ORDER BY created_at DESC`
	rows, err := r.db.Pool.Query(ctx, query, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var returnedAnomalies []Anomaly

		for rows.Next(){
		var anomaly Anomaly
		rows.Scan(
			&anomaly.AccountID,
			&anomaly.Service,
			&anomaly.ExpectedCost,
			&anomaly.ActualCost,
			&anomaly.Deviation,
			&anomaly.Severity,
			&anomaly.DetectedAt,
		)

		returnedAnomalies = append(returnedAnomalies, anomaly)
	};
		return returnedAnomalies, nil
}