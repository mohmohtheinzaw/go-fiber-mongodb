package routes

import (
	Controllers "go_jwt/controller"
	Middleware "go_jwt/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(router fiber.Router) {
	router.Post("/", Controllers.Create)
	router.Get("/", Middleware.AdminAuthMiddleware, Controllers.GetAll)
	router.Get("/:id", Controllers.GetOneAdmin)
	router.Put("/:id", Controllers.Update)
	router.Delete("/:id", Controllers.Delete)
	router.Post("/login", Controllers.LoginAdmin)
	//router.Get("/validate/test", AdminControllers.Protected)
}

func UserRoute(router fiber.Router) {
	router.Post("/", Controllers.RegisterCustomer)
	router.Post("/login", Controllers.LoginCustomer)
	router.Get("/", Middleware.UserAuthMiddleware, Controllers.GetAllUsers)
}
