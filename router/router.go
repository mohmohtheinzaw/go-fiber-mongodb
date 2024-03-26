package routes

import (
	Controllers "go_jwt/controller"
	Middleware "go_jwt/middleware"
	"go_jwt/middleware/ownership"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(router fiber.Router) {
	router.Post("/", Controllers.Create)
	router.Get("/", Middleware.AdminAuthMiddleware, Controllers.GetAll)
	router.Get("/:id", Middleware.AdminAuthMiddleware, ownership.IsCurrentUser, Controllers.GetOneAdmin)
	router.Put("/:id", Middleware.AdminAuthMiddleware, ownership.IsCurrentUser, Controllers.Update)
	router.Delete("/:id", Middleware.AdminAuthMiddleware, Controllers.Delete)
	router.Post("/login", Controllers.LoginAdmin)
	//router.Get("/validate/test", AdminControllers.Protected)
}

func UserRoute(router fiber.Router) {
	router.Post("/", Controllers.RegisterCustomer)
	router.Post("/login", Controllers.LoginCustomer)
	router.Get("/", Middleware.UserAuthMiddleware, Controllers.GetAllUsers)
}
