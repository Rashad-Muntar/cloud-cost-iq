package recommendation

import (
	"time"

	"github.com/google/uuid"
)

type Recommendation struct {
	ID string
	AccountID uuid.UUID
	Type string
	Title string
	Description string
	EstimatedMonthlySavings float64
	Severity string
	Status string
	CreatedAt time.Time
}