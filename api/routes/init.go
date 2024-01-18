package routes

import (
	"github.com/gofiber/fiber/v2"
)

// Open routes, do not require authorization
func InitRoutes(app *fiber.App) {
	initAuthRoutes(app)
	initUserRoutes(app)
	initCMSRoutes(app)
	initEventRoutes(app)
	initServicesRoutes(app)
}

// Require Authorization
func InitProtectedRoutes(app *fiber.App) {
	initProtectedAuthRoutes(app)
	initProtectedCMSRoutes(app)
	initProtectedCompanyRoutes(app)
	initProtectedEventRoutes(app)
	initConfigRoutes(app) // hidden to avoid abuse
}
