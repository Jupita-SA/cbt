package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("dashboard", fiber.Map{
			"Exams": exams,
		})
	})
	app.Get("/admin", func(c *fiber.Ctx) error {
		return c.Render("admin/dashboard", fiber.Map{
			"Exams": exams,
		})
	})

	app.Get("/analytics", func(c *fiber.Ctx) error {
		return c.Render("analytics", fiber.Map{})
	})

	app.Get("/settings", func(c *fiber.Ctx) error {
		return c.Render("settings", fiber.Map{})
	})

	manageCourseRoutes(app)
	questionRoutes(app)
	resultRoutes(app)
}
