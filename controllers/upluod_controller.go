package controllers

import (
	"begres/configs"
	"begres/models"
	"begres/responses"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var uploadCollention *mongo.Collection = configs.GetCollection(configs.DB, "upload")

//var validate = validator.New()

func CreatetoUpload(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var upload models.Upload
	defer cancel()
	file, errFile := c.FormFile("file")
	if errFile != nil {
		fmt.Println("error")
	}
	filename := file.Filename
	//fmt.Print(filename)
	errSaveFile := c.SaveFile(file, fmt.Sprint("./public/docs/", filename))

	if errSaveFile != nil {
		fmt.Println("error save file")
	}
	if err := c.BodyParser(&upload); err != nil {
		//fmt.Println("error data")
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	/*
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&upload); validationErr != nil {
			//fmt.Println("error data")
			return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
		}
	*/

	newUopload := models.Upload{
		Id:        primitive.NewObjectID(),
		File:      filename,
		Idpagu:    upload.Idpagu,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := uploadCollention.InsertOne(ctx, newUopload)
	if err != nil {

		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})

}
