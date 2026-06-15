package analytics

import "github.com/google/uuid"

type ServiceBreakdown struct {
	Service string `json:"service"`

	Cost float64 `json:"cost"`
}

type CostSummary struct {
	TotalCost float64 `json:"total_cost"`

	Breakdown  []ServiceBreakdown `json:"breakdown"`
	InternalID uuid.UUID
}