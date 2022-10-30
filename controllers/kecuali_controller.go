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

var kecualiCollection *mongo.Collection = configs.GetCollection(configs.DB, "kecuali")

func CreateKecuali(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var kecuali models.Kecuali
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&kecuali); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&kecuali); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newKecuali := models.Kecuali{
		Id:     primitive.NewObjectID(),
		Name:   kecuali.Name,
		Paket:  kecuali.Paket,
		Pagu:   kecuali.Pagu,
		Jadwal: kecuali.Jadwal,
		Idpagu: kecuali.Idpagu,
	}

	result, err := kecualiCollection.InsertOne(ctx, newKecuali)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllKecuali(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var kecualis []models.Kecuali
	defer cancel()

	results, err := kecualiCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singKecuali models.Kecuali
		if err = results.Decode(&singKecuali); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		kecualis = append(kecualis, singKecuali)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": kecualis}},
	)
}

func GetFilterKecuali(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	kecualiId := c.Params("paguId")
	var kecualis []models.Kecuali
	defer cancel()
	//objId, _ := primitive.ObjectIDFromHex(paguId)
	// csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	results, err := kecualiCollection.Find(ctx, bson.M{"idpagu": kecualiId})
	//err := paguCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&pagu)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleKecuali models.Kecuali
		if err = results.Decode(&singleKecuali); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		kecualis = append(kecualis, singleKecuali)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": kecualis}},
	)
}

func DeleteKecuali(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	kecualiId := c.Params("paguId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(kecualiId)

	result, err := kecualiCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Pagu with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Pagu successfully deleted!"}},
	)
}

func EditKecuali(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	kecualiId := c.Params("paguId")
	var kecuali models.Kecuali
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(kecualiId)

	//validate the request body
	if err := c.BodyParser(&kecuali); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&kecuali); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": kecuali.Name, "paket": kecuali.Paket, "pagu": kecuali.Pagu, "jadwal": kecuali.Jadwal}

	result, err := kecualiCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updateKecuali models.Kecuali
	if result.MatchedCount == 1 {
		err := langsungCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateKecuali)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updateKecuali}})
}

func GetKecuali(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	kecualiId := c.Params("paguId")
	var kecuali models.Kecuali
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(kecualiId)

	err := langsungCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&kecuali)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": kecuali}})
}
