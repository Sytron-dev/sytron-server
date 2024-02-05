package routes

import (
	"github.com/gofiber/fiber/v2"

	controller "sytron-server/api/handlers/service"
)

func initServicesRoutes(router *fiber.App) {
	// bed & breakfast
	router.Get("/services/bnb", controller.GetBNBs())
	router.Get("/services/bnb/:id", controller.GetSingleBNB()) // also gets assets and prices
}

func initProtectedServicesRoutes(router *fiber.App) {
	router.Post("/services/bnb", controller.CreateBNB())
	router.Post("/services/bnb/:id/image", controller.UploadBNBImage())
	router.Put("/services/bnb/:id", controller.UpdateBNB())
	router.Delete("/services/bnb/:id", controller.DeleteBNB())
}
