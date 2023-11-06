package routes

import (
	"github.com/gofiber/fiber/v2"

	controller "sytron-server/api/handlers/cms"
)

func initCMSRoutes(router *fiber.App) {
	// destinations
	router.Get("/destinations", controller.GetDestinations())
	router.Get("/destinations/:id", controller.GetSingleDestination())
}

func initProtectedCMSRoutes(router *fiber.App) {
	// destinations
	router.Post("/destinations", controller.CreateDestination())
	router.Put("/destinations/:id", controller.UpdateDestination())
	router.Put("/destinations/:id/upload/image", controller.UploadDestinationImage())
	router.Delete("/destinations/:id", controller.DeleteDestination())
}
