package jobs

import (
	"FlowForge/internal/domain/job"
	"context"
	"errors"

	"github.com/google/uuid"
)

// Get Job - Returns a single job by ID (job.ErrNotFound when it doesn't exist)
func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (job.Job, error) {
	j, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if !errors.Is(err, job.ErrNotFound) {
			s.log.ErrorContext(ctx, "get job", "id", id, "err", err)
		}
		return job.Job{}, err
	}

	return j, nil
}

// List all Jobs - Returns all jobs from the database
func (s *Service) ListAllJobs(ctx context.Context) ([]job.Job, error) {
	list, err := s.repo.ListAllJobs(ctx)
	if err != nil {
		s.log.ErrorContext(ctx, "list all jobs", "err", err)
		return nil, err
	}

	return list, nil
}
