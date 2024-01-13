package routes

import (
	"github.com/gofiber/fiber/v2"

	controller "sytron-server/api/handlers/auth"
)

func initAuthRoutes(app *fiber.App) {
	app.Post("/auth/users", controller.LoginUser)

	// backoffice
	app.Post("/auth/backofficer/login", controller.LoginBackOfficer())

	// merchant
	// router.Post("/auth/merchant/create", controller.CreateMerchantAuth)
	// router.Post("/auth/merchant/login", controller.LoginMerchant)
}

// requires user to be logged in
func initProtectedAuthRoutes(app *fiber.App) {
	// only logged in backofficers can create other backofficers
	app.Post("/auth/backofficer/create", controller.CreateBackOfficeAuth())
}
