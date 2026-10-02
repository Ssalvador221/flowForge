CREATE TABLE jobs (
    id           UUID PRIMARY KEY,
    type         TEXT        NOT NULL,
    status       TEXT        NOT NULL DEFAULT 'Pending',
    payload      JSONB       NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    started_at   TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,

    CONSTRAINT jobs_status_check CHECK (status IN ('Pending', 'Processing', 'Completed', 'Cancelled', 'Failed'))
);

CREATE INDEX idx_jobs_status_created_at ON jobs (status, created_at);

CREATE TABLE job_attempts (
    id             UUID PRIMARY KEY,
    job_id         UUID        NOT NULL REFERENCES jobs (id) ON DELETE CASCADE,
    attempt_number INT         NOT NULL,
    succeeded      BOOLEAN     NOT NULL DEFAULT false,
    error          TEXT,
    started_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    finished_at    TIMESTAMPTZ,

    CONSTRAINT job_attempts_job_attempt_unique UNIQUE (job_id, attempt_number)
);

CREATE TABLE job_events (
    id         UUID PRIMARY KEY,
    job_id     UUID        NOT NULL REFERENCES jobs (id) ON DELETE CASCADE,
    type       TEXT        NOT NULL,
    data       JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_job_events_job_id ON job_events (job_id, created_at);

CREATE TABLE job_results (
    id         UUID PRIMARY KEY,
    job_id     UUID        NOT NULL UNIQUE REFERENCES jobs (id) ON DELETE CASCADE,
    data       JSONB       NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
