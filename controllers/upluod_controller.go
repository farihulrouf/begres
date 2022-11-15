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
	"go.mongodb.org/mongo-driver/bson"
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

func GetFilterUpload(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	var uploads []models.Upload
	defer cancel()
	//objId, _ := primitive.ObjectIDFromHex(paguId)
	// csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	results, err := uploadCollention.Find(ctx, bson.M{"idpagu": paguId})
	//err := paguCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&pagu)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUpload models.Upload
		if err = results.Decode(&singleUpload); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		uploads = append(uploads, singleUpload)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": uploads}},
	)
}
