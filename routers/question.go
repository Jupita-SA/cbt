package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Jupita-SA/cbt/controllers"
)

func questionRoutes(app *fiber.App) {
	app.Get("/question", controllers.GetQuestion)
	app.Post("/question", controllers.PostQuestion)
}
