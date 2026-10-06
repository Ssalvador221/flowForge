package database

import (
	"FlowForge/internal/application/jobs"
	"FlowForge/internal/domain/job"
	"FlowForge/internal/infrastructure/database/db"
	"FlowForge/internal/utils/consts"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// JobRepository - Postgres implementation of jobs.JobRepository backed by sqlc queries
type JobRepository struct {
	q db.Querier
}

// Compile-time check that the adapter satisfies the application contract
var _ jobs.JobRepository = (*JobRepository)(nil)

func NewJobRepository(q db.Querier) *JobRepository {
	return &JobRepository{q: q}
}

func (r *JobRepository) GetByID(ctx context.Context, id uuid.UUID) (job.Job, error) {
	row, err := r.q.GetJob(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return job.Job{}, job.ErrNotFound
	}
	if err != nil {
		return job.Job{}, fmt.Errorf("get job %s: %w", id, err)
	}

	return toDomainJob(row), nil
}

func (r *JobRepository) ListAllJobs(ctx context.Context) ([]job.Job, error) {
	rows, err := r.q.ListJobs(ctx)
	if err != nil {
		return nil, fmt.Errorf("list jobs: %w", err)
	}

	jobs := make([]job.Job, 0, len(rows))
	for _, row := range rows {
		jobs = append(jobs, toDomainJob(row))
	}

	return jobs, nil
}

func (r *JobRepository) Create(ctx context.Context, j job.Job) (job.Job, error) {
	row, err := r.q.CreateJob(ctx, db.CreateJobParams{
		ID:      j.ID,
		Type:    string(j.Type),
		Payload: []byte(j.Payload),
	})
	if err != nil {
		return job.Job{}, fmt.Errorf("create job: %w", err)
	}

	return toDomainJob(row), nil
}

// Maps the sqlc row model to the domain entity
func toDomainJob(row db.Job) job.Job {
	return job.Job{
		ID:          row.ID,
		Payload:     string(row.Payload),
		Type:        consts.JobType(row.Type),
		Status:      consts.JobStatus(row.Status),
		CreatedAt:   row.CreatedAt,
		StartedAt:   row.StartedAt,
		CompletedAt: row.CompletedAt,
	}
}
