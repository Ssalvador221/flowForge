package job

import (
	"github.com/google/uuid"
	"time"
)

type JobAttempt struct {
	ID            uuid.UUID
	JobID         uuid.UUID
	AttemptNumber int
	Succeded      bool
	Error         *string
	StartedAt     time.Time
	FineshedAt    *time.Time
}
