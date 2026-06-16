package anomaly

import "time"

type Anomaly struct {
	ID string
	AccountID string
	Service string
	ExpectedCost float64
	ActualCost float64
	Deviation float64
	Severity string
	DetectedAt time.Time
}