package http

import (
	"FlowForge/internal/transport/http/handlers"
	"FlowForge/internal/transport/http/routes"

	"github.com/gofiber/fiber/v3"
)

type Handlers struct {
	Jobs *handlers.JobHandler
}

// Router Group
func RegisterRoutes(app *fiber.App, h Handlers) {
	api := app.Group("/api")

	routes.RegisterJobRoutes(api, h.Jobs)
	routes.RegisterEmailRoutes(api)
	routes.RegisterReportRoutes(api)
}
