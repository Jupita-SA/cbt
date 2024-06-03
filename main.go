package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/Jupita-SA/cbt/routes"
	"github.com/Jupita-SA/cbt/utils"
)

func main() {
	utils.InitDB()
	defer utils.CloseDB()

	engine := html.New("./views", ".html")

	engine.AddFunc("add", func(a, b int) int {
		return a + b
	})

	engine.AddFunc("sub", func(a, b int) int {
		return a - b
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	// Setup routes
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
