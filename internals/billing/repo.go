package billing

import (
	"context"
	"strconv"
	"time"

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

// internal/api/handlers_account.go
// Add this handler


// // Optional: Add PATCH for partial updates (more RESTful)
// func (h *Handlers) PatchAccount(w http.ResponseWriter, r *http.Request) {
//     // Same as UpdateAccount but only updates provided fields
//     // This is more RESTful than PUT for partial updates
    
//     idStr := chi.URLParam(r, "id")
//     id, err := uuid.Parse(idStr)
//     if err != nil {
//         http.Error(w, "Invalid account ID format", http.StatusBadRequest)
//         return
//     }
    
//     var input account.UpdateAccountInput
//     if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
//         http.Error(w, "Invalid request body", http.StatusBadRequest)
//         return
//     }
    
//     input.InternalID = id
//     updated, err := h.accountService.UpdateAccount(r.Context(), input)
//     if err != nil {
//         if strings.Contains(err.Error(), "not found") {
//             http.Error(w, err.Error(), http.StatusNotFound)
//         } else {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//         }
//         return
//     }
    
//     writeJSON(w, http.StatusOK, updated)
// }

