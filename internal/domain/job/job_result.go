package job

import (
	"time"

	"github.com/google/uuid"
)

type JobResult struct {
	ID        uuid.UUID
	JobID     uuid.UUID
	Data      string
	CreatedAt time.Time
}
