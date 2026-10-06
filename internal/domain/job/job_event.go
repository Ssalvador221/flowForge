package job

import (
	"FlowForge/internal/utils/consts"
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
