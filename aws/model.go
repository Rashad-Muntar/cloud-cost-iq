package aws
import(
	"time"
)
type AWSConnection struct {
	ID string

	AccountID string

	RoleARN string

	ExternalID string

	Status string

	LastSyncAt *time.Time
}

const (
	ConnectionPending = "pending"
	ConnectionActive = "active"
	ConnectionFailed = "failed"
)