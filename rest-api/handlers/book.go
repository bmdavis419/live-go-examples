package handlers

import (
	"context"
	"fmt"

	"github.com/bmdavis419/live-go-examples/rest-api/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type newBookDTO struct {
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	ISBN      string `json:"isbn" bson:"isbn"`
	LibraryId string `json:"libraryId" bson:"libraryId"`
}

// THIS IS BEING DUMB, I THINK ITS A ME PROBLEM BUT I RAN OUT OF TIME :)
func CreateBook(c *fiber.Ctx) error {
	createData := new(newBookDTO)
	if err := c.BodyParser(createData); err != nil {
		return err
	}

	// get the collection reference
	coll := database.GetCollection("libraries")

	// update one library
	res, err := coll.UpdateOne(context.TODO(), bson.D{{Key: "name", Value: "second library"}}, bson.D{{Key: "$push", Value: bson.M{"books": bson.M{"test": "test"}}}})

	// get the filter
	filter := bson.D{{Key: "id", Value: createData.LibraryId}}
	fmt.Println(filter)
	// nBookData := models.Book{
	// 	Title:  createData.Title,
	// 	Author: createData.Author,
	// 	ISBN:   createData.ISBN,
	// }
	// updatePayload := bson.D{{Key: "$push", Value: bson.M{"books": bson.M{"test": "test"}}}}

	// update the library
	if err != nil {
		return err
	}

	fmt.Println(res.ModifiedCount)

	return c.SendString("Book created successfully")
}
