package jobs

import (
	"FlowForge/internal/domain/job"
	"context"

	"github.com/google/uuid"
)

// JobRepository - Storage contract required by the jobs use cases
type JobRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (job.Job, error)
	ListAllJobs(ctx context.Context) ([]job.Job, error)
	Create(ctx context.Context, j job.Job) (job.Job, error)
}
