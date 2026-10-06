package job

import (
	"FlowForge/internal/utils/consts"
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID          uuid.UUID
	Payload     string
	Type        consts.JobType
	Status      consts.JobStatus
	CreatedAt   time.Time
	StartedAt   *time.Time
	CompletedAt *time.Time
}
