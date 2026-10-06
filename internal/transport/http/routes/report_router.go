package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func RegisterReportRoutes(router fiber.Router) {
	report := router.Group("/report")

	// TOOD: Remove this return and add the real func
	report.Get("/", func(c fiber.Ctx) error {
		return c.SendString("report 1")
	})

	report.Get("/:id", func(c fiber.Ctx) error {
		return c.SendString("report by ID")
	})

	report.Post("/", func(c fiber.Ctx) error {
		var reports []string
		if err := c.Bind().JSON(&reports); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
		}
		return c.SendString(strconv.Itoa(len(reports)))
	})
}
