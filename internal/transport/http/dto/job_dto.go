package dto

import (
	"FlowForge/internal/domain/job"
	"FlowForge/internal/utils/consts"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Job Response DTO
type JobResponseDTO struct {
	ID          uuid.UUID        `json:"id"`
	Payload     json.RawMessage  `json:"payload"`
	Status      consts.JobStatus `json:"status"`
	Type        consts.JobType   `json:"type"`
	CreatedAt   time.Time        `json:"created_at"`
	StartedAt   *time.Time       `json:"started_at"`
	CompletedAt *time.Time       `json:"completed_at"`
}

// Maps a domain job to its HTTP response
func ToJobResponse(j job.Job) JobResponseDTO {
	return JobResponseDTO{
		ID:          j.ID,
		Payload:     json.RawMessage(j.Payload),
		Status:      j.Status,
		Type:        j.Type,
		CreatedAt:   j.CreatedAt,
		StartedAt:   j.StartedAt,
		CompletedAt: j.CompletedAt,
	}
}

// Maps a list of domain jobs; never nil, so it always encodes as [] instead of null
func ToJobResponseList(jobs []job.Job) []JobResponseDTO {
	out := make([]JobResponseDTO, 0, len(jobs))
	for _, j := range jobs {
		out = append(out, ToJobResponse(j))
	}
	return out
}

// Job Request DTO - payload is any JSON object, stored as-is in the JSONB column
type JobRequestDTO struct {
	Payload json.RawMessage `json:"payload"`
	Type    string          `json:"type"`
}
