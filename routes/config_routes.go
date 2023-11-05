package routes

import (
	"github.com/gofiber/fiber/v2"

	controller "sytron-server/controllers/configs_controller"
)

func initConfigRoutes(router *fiber.App) {
	router.Get("/configs/countries", controller.GetCountries())
	router.Get("/configs/cities", controller.GetCities())
}
