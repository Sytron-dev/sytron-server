package routes

import (
	"github.com/gofiber/fiber/v2"

	controller "sytron-server/api/handlers/auth"
)

func initAuthRoutes(router *fiber.App) {
	router.Post("/auth/users", controller.LoginUser)

	// backoffice
	router.Post("/auth/backofficer/login", controller.LoginBackOfficer)

	// merchant
	router.Post("/auth/merchant/create", controller.CreateMerchantAuth)
	router.Post("/auth/merchant/login", controller.LoginMerchant)
}

// requires user to be logged in
func initProtectedAuthRoutes(router *fiber.App) {
	// only logged in backofficers can create other backofficers
	router.Post("/auth/backofficer/create", controller.CreateBackOfficeAuth)
}
