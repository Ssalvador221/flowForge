package handlers

import (
	"FlowForge/internal/application/jobs"
	"FlowForge/internal/domain/job"
	"FlowForge/internal/transport/http/dto"
	"FlowForge/internal/utils/consts"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// JobHandler - HTTP controller for /jobs; unexpected errors are logged by the server's errorHandler
type JobHandler struct {
	svc *jobs.Service
}

func NewJobHandler(svc *jobs.Service) *JobHandler {
	return &JobHandler{svc: svc}
}

// GET /jobs
func (h *JobHandler) ListAll(c fiber.Ctx) error {
	list, err := h.svc.ListAllJobs(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(dto.ToJobResponseList(list))
}

// POST /jobs
func (h *JobHandler) Create(c fiber.Ctx) error {
	var req dto.JobRequestDTO
	if err := c.Bind().JSON(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}

	j, err := h.svc.Create(c.Context(), jobs.CreateJobInput{
		Type:    consts.JobType(req.Type),
		Payload: req.Payload,
	})
	switch {
	case errors.Is(err, job.ErrInvalidType), errors.Is(err, job.ErrInvalidPayload):
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	case err != nil:
		return err
	}

	c.Location("/api/jobs/" + j.ID.String())
	return c.Status(fiber.StatusCreated).JSON(dto.ToJobResponse(j))
}

// GET /jobs/:id
func (h *JobHandler) GetJobByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid job id")
	}

	j, err := h.svc.GetByID(c.Context(), id)
	if errors.Is(err, job.ErrNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "job not found")
	}
	if err != nil {
		return err
	}

	return c.JSON(dto.ToJobResponse(j))
}
