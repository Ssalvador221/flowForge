package routes

import (
	"FlowForge/internal/transport/http/handlers"

	"github.com/gofiber/fiber/v3"
)

// Jobs Router Group
func RegisterJobRoutes(router fiber.Router, h *handlers.JobHandler) {
	jobs := router.Group("/jobs")

	jobs.Get("/", h.ListAll)
	jobs.Get("/:id", h.GetJobByID)
	jobs.Post("/", h.Create)
}
