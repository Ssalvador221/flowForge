package http

import (
	handlers "FlowForge/internal/transport/http/routes"

	"github.com/gofiber/fiber/v3"
)

// Router Group
func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	handlers.RegisterJobRoutes(api)
	handlers.RegisterEmailRoutes(api)
	handlers.RegisterReportRoutes(api)
}
