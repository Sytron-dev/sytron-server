package routes

import (
	"github.com/gofiber/fiber/v2"

	controller "sytron-server/api/handlers"
)

// UserRoutes function
func initUserRoutes(router *fiber.App) {
	router.Post("/users/signup", controller.SignUp())
	router.Post("/users/login", controller.Login())
}
