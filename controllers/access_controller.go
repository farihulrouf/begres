package controllers

import (
	"begres/configs"
	"begres/models"
	"begres/responses"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var accessCollection *mongo.Collection = configs.GetCollection(configs.DB, "access")

//var validation = validator.New()

func CreateAccess(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var access models.Access
	defer cancel()

	//validation the request body
	if err := c.BodyParser(&access); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validation required fields
	if validationErr := validation.Struct(&access); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}
	/*
		Name:   link.Name,
			Link:   link.Link,
			Idpagu: link.Idpagu,
	*/

	newAccess := models.Access{
		Id:     primitive.NewObjectID(),
		Name:   access.Name,
		Idpagu: access.Idpagu,
		UserId: access.UserId,
	}

	result, err := linkCollection.InsertOne(ctx, newAccess)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}
