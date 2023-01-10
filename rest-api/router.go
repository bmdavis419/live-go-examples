package main

import (
	"github.com/bmdavis419/live-go-examples/rest-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func generateApp() *fiber.App {
	app := fiber.New()

	// create health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// create the library group and routes
	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.TestHandler)

	return app
}
