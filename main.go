package main

import (
	routes "go_jwt/router"
	"go_jwt/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	app.Listen(":6001")

}

func setupRoutes(app *fiber.App) {
	// start database
	database.StartMongoDB()

	// // defer closing database
	// defer database.CloseMongoDB()

	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	// api group
	api := app.Group("/api")
	routes.AdminRoute(api.Group("/admins"))
	//routes.AuthRoute(api.Group("/auth"))
}
