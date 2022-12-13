package controllers

import (
	"begres/configs"
	"begres/models"
	"begres/responses"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var linkCollection *mongo.Collection = configs.GetCollection(configs.DB, "links")

//var validation = validator.New()

func CreateLink(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var link models.Link
	defer cancel()

	//validation the request body
	if err := c.BodyParser(&link); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validation required fields
	if validationErr := validation.Struct(&link); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newLink := models.Link{
		Id:     primitive.NewObjectID(),
		Name:   link.Name,
		Link:   link.Link,
		Idpagu: link.Idpagu,
	}

	result, err := linkCollection.InsertOne(ctx, newLink)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllLink(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var links []models.Link
	defer cancel()

	results, err := linkCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleLink models.Link
		if err = results.Decode(&singleLink); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		links = append(links, singleLink)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": links}},
	)
}

func DeleteLink(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(paguId)

	result, err := linkCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.Response{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Pagu with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Pagu successfully deleted!"}},
	)
}

func EditLink(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	var link models.Link
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(paguId)

	//validate the request body
	if err := c.BodyParser(&link); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&link); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": link.Name, "link": link.Link, "idpagu": link.Idpagu}

	result, err := linkCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updatePagu models.Pagu
	if result.MatchedCount == 1 {
		err := paguCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatePagu)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatePagu}})
}

func GetLink(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	var link models.Link
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(paguId)

	err := linkCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&link)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": link}})
}
