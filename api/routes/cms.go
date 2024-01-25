package routes

import (
	"github.com/gofiber/fiber/v2"

	controller "sytron-server/api/handlers/cms"
)

func initCMSRoutes(router *fiber.App) {
	// destinations
	router.Get("/cms/destinations", controller.GetDestinations())
	router.Get("/cms/destinations/:id", controller.GetSingleDestination())
}

func initProtectedCMSRoutes(router *fiber.App) {
	// destinations
	router.Post("/cms/destinations", controller.CreateDestination())
	router.Put("/cms/destinations/:id", controller.UpdateDestination())
	router.Put("/cms/destinations/:id/upload/image", controller.UploadDestinationImage())
	router.Delete("/cms/destinations/:id", controller.DeleteDestination())
}
