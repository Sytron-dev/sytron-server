package routes

import (
	controller "sytron-server/controllers/company_controllers"

	"github.com/gofiber/fiber/v2"
)

func initProtectedCompanyRoutes(router *fiber.App) {
	router.Get("/companies", controller.GetCompanies())
	router.Get("/companies/:id", controller.GetSingleCompany())
	router.Post("/companies", controller.CreateCompany())
	router.Put("/companies/:id", controller.UpdateCompany())
}
