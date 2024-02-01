package routes

import (
	AdminControllers "go_jwt/controller"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(router fiber.Router) {
	router.Post("/", AdminControllers.Create)
	router.Get("/", AdminControllers.GetAll)
	router.Get("/:id", AdminControllers.GetOneAdmin)
	router.Put("/:id", AdminControllers.Update)
	router.Delete("/:id", AdminControllers.Delete)
	router.Post("/login", AdminControllers.LoginAdmin)
	router.Get("/validate/test", AdminControllers.Validate)
}
