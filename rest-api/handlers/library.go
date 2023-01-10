package handlers

import (
	"context"

	"github.com/bmdavis419/live-go-examples/rest-api/database"
	"github.com/bmdavis419/live-go-examples/rest-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// DELETE
func DeleteLibrary(c *fiber.Ctx) error {
	// get the id from the params
	id := c.Params("id")

	libraryCollection := database.GetCollection("libraries")
	_, err := libraryCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		return err
	}

	return c.SendString("Library deleted successfully")
}

// GET
func GetLibraries(c *fiber.Ctx) error {
	libraryCollection := database.GetCollection("libraries")
	cursor, err := libraryCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return err
	}

	var libraries []models.Library
	if err = cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}

	return c.JSON(libraries)
}

type libraryDTO struct {
	Name    string   `json:"name" bson:"name"`
	Address string   `json:"address" bson:"address"`
	Empty   []string `json:"no_exists" bson:"books"`
}

// POST
func CreateLibrary(c *fiber.Ctx) error {
	nLibrary := new(libraryDTO)

	if err := c.BodyParser(nLibrary); err != nil {
		return err
	}

	// MongoDB being questionable
	nLibrary.Empty = make([]string, 0)

	libraryCollection := database.GetCollection("libraries")
	nDoc, err := libraryCollection.InsertOne(context.TODO(), nLibrary)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"id": nDoc.InsertedID})
}
