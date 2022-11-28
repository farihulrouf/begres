package controllers

import (
	"begres/configs"
	"begres/models"
	"begres/responses"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var paguCollection *mongo.Collection = configs.GetCollection(configs.DB, "pagus")
var validation = validator.New()

func CreatePagu(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var pagu models.Pagu
	defer cancel()

	//validation the request body
	if err := c.BodyParser(&pagu); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validation required fields
	if validationErr := validation.Struct(&pagu); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newPagu := models.Pagu{
		Id:        primitive.NewObjectID(),
		Name:      pagu.Name,
		Paguopdp:  pagu.Paguopdp,
		Paguorp:   pagu.Paguorp,
		Filetipe:  pagu.Filetipe,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := paguCollection.InsertOne(ctx, newPagu)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllPagu(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var pagus []models.Pagu
	defer cancel()

	results, err := paguCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singlePagu models.Pagu
		if err = results.Decode(&singlePagu); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		pagus = append(pagus, singlePagu)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": pagus}},
	)
}

func GetAllFilter(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	findOptions := options.Find()
	filter := bson.M{}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	var perPage int64 = 2

	total, _ := paguCollection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page) - 1) * perPage)
	findOptions.SetLimit(perPage)
	var pagus []models.Pagu
	defer cancel()

	results, err := paguCollection.Find(ctx, filter, findOptions)
	fmt.Print(total)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singlePagu models.Pagu
		if err = results.Decode(&singlePagu); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		pagus = append(pagus, singlePagu)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": pagus}},
	)
}

func DeletePagu(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(paguId)

	result, err := paguCollection.DeleteOne(ctx, bson.M{"id": objId})
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

func EditPagu(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	var pagu models.Pagu
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(paguId)

	//validate the request body
	if err := c.BodyParser(&pagu); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&pagu); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": pagu.Name, "paguopdp": pagu.Paguopdp, "paguorp": pagu.Paguorp, "updatedat": time.Now()}

	result, err := paguCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
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

func GetPagu(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	var pagu models.Pagu
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(paguId)

	err := paguCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&pagu)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": pagu}})
}

func EditPaguUpload(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	var pagu models.Pagu
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(paguId)
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
	//validate the request body
	if err := c.BodyParser(&pagu); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	update := bson.M{"name": pagu.Name, "paguopdp": pagu.Paguopdp, "paguorp": pagu.Paguorp, "filetipe": filename, "updatedat": time.Now()}

	result, err := paguCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
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
