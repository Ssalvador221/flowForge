package job

import (
	"FlowForge/internal/domain/consts"
	"time"

	"github.com/google/uuid"
)

type JobEvent struct {
	ID        uuid.UUID
	JobId     uuid.UUID
	Type      consts.JobType
	Data      *string
	CreatedAt time.Time
}
