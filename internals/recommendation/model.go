package recommendation

import "time"

type Recommendation struct {
	ID string

	AccountID string

	Type string

	Title string

	Description string

	EstimatedMonthlySavings float64

	Severity string

	Status string

	CreatedAt time.Time
}