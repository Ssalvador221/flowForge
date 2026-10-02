package handlers

import "github.com/gofiber/fiber/v3"

// Jobs Router Group
func RegisterJobRoutes(router fiber.Router) {
	jobs := router.Group("/jobs")

	// TOOD: Remove this return and add the real func
	jobs.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Jobs 1")
	})

	jobs.Get("/:id", func(c fiber.Ctx) error {
		return c.SendString("Job by ID")
	})

	jobs.Post("/", func(c fiber.Ctx) error {
		var job []string
		if err := c.Bind().JSON(&job); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
		}
		if len(job) == 0 {
			return fiber.NewError(fiber.StatusBadRequest, "body must be a non-empty JSON array")
		}
		return c.SendString(job[0])
	})

}
