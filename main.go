package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"sytron-server/api/routes"
	"sytron-server/constants"
)

func main() {
	// Server

	port := "8000"
	if constants.PORT != "" {
		port = constants.PORT
	}

	app := fiber.New()
	app.Use(cors.New())

	if constants.ENVIRONMENT == "development" {
		app.Use(func(c *fiber.Ctx) error {
			log.Printf("Fetching %v\n", c.Request().URI())
			return c.Next()
		})
	}

	routes.InitRoutes(app)
	// TODO add endpoint-level protection
	routes.InitProtectedRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
