package jobs

import (
	"FlowForge/internal/domain/job"
	"FlowForge/internal/utils/consts"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// CreateJobInput - Data required to enqueue a new job
type CreateJobInput struct {
	Type    consts.JobType
	Payload json.RawMessage
}

// Create Job - Validates the input and enqueues a new Pending job
func (s *Service) Create(ctx context.Context, in CreateJobInput) (job.Job, error) {
	if !in.Type.IsValid() {
		return job.Job{}, job.ErrInvalidType
	}

	// The payload column is JSONB and workers expect an object, not a scalar/array
	var obj map[string]any
	if err := json.Unmarshal(in.Payload, &obj); err != nil || obj == nil {
		return job.Job{}, job.ErrInvalidPayload
	}

	// UUIDv7 is time-ordered, which keeps the primary key index append-only
	id, err := uuid.NewV7()
	if err != nil {
		return job.Job{}, fmt.Errorf("generate job id: %w", err)
	}

	j, err := s.repo.Create(ctx, job.Job{
		ID:      id,
		Type:    in.Type,
		Payload: string(in.Payload),
	})
	if err != nil {
		s.log.ErrorContext(ctx, "create job", "type", in.Type, "err", err)
		return job.Job{}, err
	}

	s.log.InfoContext(ctx, "job created", "id", j.ID, "type", j.Type)
	return j, nil
}
