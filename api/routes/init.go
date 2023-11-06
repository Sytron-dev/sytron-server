package routes

import (
	"github.com/gofiber/fiber/v2"
)

// Open routes, do not require authorization
func InitRoutes(router *fiber.App) {
	initAuthRoutes(router)
	initUserRoutes(router)
	initConfigRoutes(router)
	initCMSRoutes(router)
	initEventRoutes(router)
}

// Require Authorization
func InitProtectedRoutes(router *fiber.App) {
	initProtectedAuthRoutes(router)
	initProtectedCMSRoutes(router)
	initProtectedCompanyRoutes(router)
	initProtectedEventRoutes(router)
}
