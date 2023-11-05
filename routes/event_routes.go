package routes

import (
	controller "sytron-server/controllers/event_controllers"

	"github.com/gofiber/fiber/v2"
)

func initEventRoutes(router *fiber.App) {
	router.Get("/events/:id", controller.GetSingleEvent())
	router.Get("/events", controller.GetEvents())
}

func initProtectedEventRoutes(router *fiber.App) {
	router.Post("/events", controller.CreateEvent())
	router.Post("/events/:id/image", controller.UploadEventHeroImage())
	router.Put("/events/:id", controller.UpdateEvent())
	router.Delete("/events/:id", controller.DeleteEvent())
}
