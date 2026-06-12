package billing

import (
	"context"
	"strconv"
	"time"
	"fmt"
	"github.com/cloud-cost-iq/internals/db"
)

type repository struct {
	db *db.Database
}

func NewRepository(database *db.Database) Repository {  // ← returns pointer to private struct
    return &repository{db: database}
}

func (repo *repository) InsertCost(ctx context.Context, input CostRecord) error {
	query := `
		INSERT INTO cost_events (
			id, account_id, service, region,
			cost_amount, usage_amount,
			currency, usage_date, created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`
	_, err := repo.db.Pool.Exec(ctx, query,
		input.InternalID,
		input.AwsAccountID,
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
	accountID string,
	service string,
) (*DailyCostSummary, error) {
	query := `
		SELECT service, COALESCE(SUM(cost_amount),0) FROM cost_events
		WHERE usage_date >= $1
	  	AND usage_date < $2
	`
	start := date.Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)
	args := []any{start, end}

if accountID != "" {
	query += " AND account_id = $3"
	args = append(args, accountID)
}

if service != "" {
	query += " AND service = $" + strconv.Itoa(len(args)+1)
	args = append(args, service)
}

query += " GROUP BY service ORDER BY SUM(cost_amount) DESC"
	rows, _ := repo.db.Pool.Query(ctx, query, args...)
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

func (r *repository) GetCostsWithAccountFilter(ctx context.Context, filter AccountFilter) ([]CostRecordWithAccount, error) {
    query := `
        SELECT 
            ce.id, ce.account_id, ce.service, ce.region,
            ce.cost_amount, ce.usage_amount, ce.currency,
            ce.usage_date, ce.created_at,
            a.name as account_name,
            a.environment,
            a.aws_account_id
        FROM cost_events ce
        INNER JOIN accounts a ON ce.account_id = a.id
        WHERE a.deleted_at IS NULL
    `
    
    args := []interface{}{}
    argPos := 1
    
    if filter.Environment != nil {
        query += fmt.Sprintf(" AND a.environment = $%d", argPos)
        args = append(args, *filter.Environment)
        argPos++
    }
    
    if filter.IsActive != nil {
        query += fmt.Sprintf(" AND a.is_active = $%d", argPos)
        args = append(args, *filter.IsActive)
        argPos++
    }
    
    if filter.StartDate != nil {
        query += fmt.Sprintf(" AND ce.usage_date >= $%d", argPos)
        args = append(args, *filter.StartDate)
        argPos++
    }
    
    if filter.EndDate != nil {
        query += fmt.Sprintf(" AND ce.usage_date <= $%d", argPos)
        args = append(args, *filter.EndDate)
        argPos++
    }
    
    query += " ORDER BY ce.usage_date DESC"
    
    rows, err := r.db.Pool.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var results []CostRecordWithAccount
    for rows.Next() {
        var result CostRecordWithAccount
        err := rows.Scan(
            &result.InternalID, &result.AwsAccountID, &result.Service, &result.Region,
            &result.CostAmount, &result.UsageAmount, &result.Currency,
            &result.UsageDate, &result.CreatedAt,
            &result.AccountName, &result.Environment, &result.AWSAccountID,
        )
        if err != nil {
            return nil, err
        }
        results = append(results, result)
    }
    
    return results, nil
}


