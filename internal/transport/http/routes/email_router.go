package routes

import "github.com/gofiber/fiber/v3"

func RegisterEmailRoutes(router fiber.Router) {
	email := router.Group("/email")

	// TOOD: Remove this return and add the real func
	email.Get("/", func(c fiber.Ctx) error {
		return c.SendString("email 1")
	})

	email.Get("/:id", func(c fiber.Ctx) error {
		return c.SendString("Email by ID")
	})

	email.Post("/", func(c fiber.Ctx) error {
		var emails []string
		if err := c.Bind().JSON(&emails); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
		}
		if len(emails) == 0 {
			return fiber.NewError(fiber.StatusBadRequest, "body must be a non-empty JSON array")
		}
		return c.SendString(emails[0])
	})
}
