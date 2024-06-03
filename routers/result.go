package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Jupita-SA/cbt/controllers"
)

func resultRoutes(app *fiber.App) {
	app.Get("/result", controllers.GetResult)
}
