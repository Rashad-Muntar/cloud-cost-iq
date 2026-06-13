package analytics

import "time"

type CostQuery struct {
	AwsAccountID string

	From time.Time

	To time.Time
}