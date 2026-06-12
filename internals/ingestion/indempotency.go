package ingestion

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateIdempotencyKey(r RawCostRecord) string {

	raw := fmt.Sprintf(
		"%s:%s:%s:%s:%f:%f:%d",
		r.AwsAccountID,
		r.Service,
		r.Region,
		r.Currency,
		r.CostAmount,
		r.UsageAmount,
		r.UsageEnd.Unix(),
	)

	hash := sha256.Sum256([]byte(raw))

	return hex.EncodeToString(hash[:])
}