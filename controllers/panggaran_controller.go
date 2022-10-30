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

var panggaranCollection *mongo.Collection = configs.GetCollection(configs.DB, "panggaran")

func CreatePanggaran(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var panggaran models.Panggaran
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&panggaran); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&panggaran); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newPanggaran := models.Panggaran{
		Id:     primitive.NewObjectID(),
		Name:   panggaran.Name,
		Jumlah: panggaran.Jumlah,
		Idpagu: panggaran.Idpagu,
	}

	result, err := panggaranCollection.InsertOne(ctx, newPanggaran)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllAnggaran(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var panggarans []models.Panggaran
	defer cancel()

	results, err := panggaranCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singPanggaran models.Panggaran
		if err = results.Decode(&singPanggaran); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		panggarans = append(panggarans, singPanggaran)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": panggarans}},
	)
}
func GetFilterAnggaran(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	var panggarans []models.Panggaran
	defer cancel()
	//objId, _ := primitive.ObjectIDFromHex(paguId)
	// csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	results, err := panggaranCollection.Find(ctx, bson.M{"idpagu": paguId})
	//err := paguCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&pagu)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singPanggaran models.Panggaran
		if err = results.Decode(&singPanggaran); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		panggarans = append(panggarans, singPanggaran)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": panggarans}},
	)
}
