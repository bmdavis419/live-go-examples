package main

import (
	"os"

	"github.com/bmdavis419/live-go-examples/rest-api/database"
	"github.com/bmdavis419/live-go-examples/rest-api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// init app
	err := initApp()
	if err != nil {
		panic(err)
	}

	// defer close database
	defer database.CloseMongoDB()

	app := generateApp()

	// get the port from the env
	port := os.Getenv("PORT")

	app.Listen(":" + port)
}

func generateApp() *fiber.App {
	app := fiber.New()

	// create health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// create the library group and routes
	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.GetLibraries)
	libGroup.Post("/", handlers.CreateLibrary)
	libGroup.Delete("/:id", handlers.DeleteLibrary)

	// :(
	libGroup.Post("/book", handlers.CreateBook)

	return app
}

func initApp() error {
	// setup env
	err := loadENV()
	if err != nil {
		return err
	}

	// setup database
	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	return nil
}

func loadENV() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
