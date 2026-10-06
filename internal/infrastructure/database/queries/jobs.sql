-- name: CreateJob :one
INSERT INTO jobs (id, type, payload)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetJob :one
SELECT * FROM jobs
WHERE id = $1;

-- name: ListJobs :many
SELECT * FROM jobs
ORDER BY created_at DESC;

-- name: ListJobsByStatus :many
SELECT * FROM jobs
WHERE status = $1
ORDER BY created_at
LIMIT $2;

-- Atomically claims the oldest pending job; SKIP LOCKED lets several workers poll concurrently.
-- name: ClaimNextPendingJob :one
UPDATE jobs
SET status = 'Processing', started_at = now()
WHERE id = (
    SELECT id FROM jobs
    WHERE status = 'Pending'
    ORDER BY created_at
    LIMIT 1
    FOR UPDATE SKIP LOCKED
)
RETURNING *;

-- name: UpdateJobStatus :one
UPDATE jobs
SET status = sqlc.arg(status)::text,
    completed_at = CASE WHEN sqlc.arg(status)::text IN ('Completed', 'Cancelled', 'Failed') THEN now() ELSE completed_at END
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: CreateJobAttempt :one
INSERT INTO job_attempts (id, job_id, attempt_number)
VALUES ($1, $2, $3)
RETURNING *;

-- name: FinishJobAttempt :one
UPDATE job_attempts
SET succeeded = $2, error = $3, finished_at = now()
WHERE id = $1
RETURNING *;

-- name: ListJobAttempts :many
SELECT * FROM job_attempts
WHERE job_id = $1
ORDER BY attempt_number;

-- name: CreateJobEvent :one
INSERT INTO job_events (id, job_id, type, data)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListJobEvents :many
SELECT * FROM job_events
WHERE job_id = $1
ORDER BY created_at;

-- name: CreateJobResult :one
INSERT INTO job_results (id, job_id, data)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetJobResult :one
SELECT * FROM job_results
WHERE job_id = $1;
