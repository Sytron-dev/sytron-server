package routes

import (
	"github.com/gofiber/fiber/v2"

	controller "sytron-server/api/handlers/event"
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
