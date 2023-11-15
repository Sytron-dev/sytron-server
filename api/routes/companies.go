package routes

import (
	"github.com/gofiber/fiber/v2"

	controller "sytron-server/api/handlers/company"
)

func initProtectedCompanyRoutes(router *fiber.App) {
	router.Get("/companies", controller.GetCompanies())
	router.Get("/companies/:id", controller.GetSingleCompany())
	router.Post("/companies", controller.CreateCompany())
	router.Put("/companies/:id", controller.UpdateCompany())
}
