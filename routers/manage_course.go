package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Jupita-SA/cbt/controllers"
)

func manageCourseRoutes(app *fiber.App) {
	app.Get("/manage_course", controllers.GetCourses)
	app.Post("/manage_course/add", controllers.AddCourse)
	app.Post("/manage_course/edit/:id", controllers.EditCourse)
	app.Post("/manage_course/delete/:id", controllers.DeleteCourse)
}
