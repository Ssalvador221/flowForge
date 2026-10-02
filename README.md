# FlowForge

FlowForge is a backend project built with Go for learning and experimenting with asynchronous job processing, background workers, and observability.

The main idea is simple: an API receives jobs, stores them, and workers process those jobs asynchronously while the application tracks their status, results, attempts, and events.

## Overview

FlowForge is designed around a job processing pipeline:

Client → API → Job → Queue/Storage → Worker → Processor → Result

A client creates a job through the API. The job is persisted with a `Pending` status and later picked up by a background worker.

The worker identifies the job type, executes the appropriate processor, and updates the job status and result.

## Architecture

The project is organized into independent components:

```text
flowforge/
│
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
│
├── internal/
│   ├── domain/
│   ├── application/
│   ├── infrastructure/
│   ├── transport/
│   ├── worker/
│   └── observability/
│
├── migrations/
├── go.mod
└── go.sum
