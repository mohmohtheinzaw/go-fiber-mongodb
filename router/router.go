package routes

import (
	Controllers "go_jwt/controller"
	Middleware "go_jwt/middleware"
	"go_jwt/middleware/ownership"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(router fiber.Router) {
	router.Post("/", Controllers.Create)
	router.Get("/", Middleware.IsAdmin, Controllers.GetAll)
	router.Get("/:id", Middleware.IsAdmin, ownership.IsCurrentUser, Controllers.GetOneAdmin)
	router.Put("/:id", Middleware.IsAdmin, ownership.IsCurrentUser, Controllers.Update)
	router.Delete("/:id", Middleware.IsAdmin, Controllers.Delete)
	router.Post("/login", Controllers.LoginAdmin)
}

func UserRoute(router fiber.Router) {
	router.Post("/", Controllers.RegisterCustomer)
	router.Post("/login", Controllers.LoginCustomer)
	router.Get("/", Middleware.IsUser, ownership.IsCurrentUser, Controllers.GetAllUsers)
}
