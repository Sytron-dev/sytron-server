package routes

import (
	"github.com/gofiber/fiber/v2"

	controller "sytron-server/api/handlers/config"
)

func initConfigRoutes(router *fiber.App) {
	router.Get("/configs/countries", controller.GetCountries())
	router.Get("/configs/cities", controller.GetCities())
}
