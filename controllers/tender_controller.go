package controllers

import (
	"begres/configs"
	"begres/models"
	"begres/responses"
	"context"

	//"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var tenderCollection *mongo.Collection = configs.GetCollection(configs.DB, "tender")

func CreateTender(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var tender models.Tender
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&tender); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&tender); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newTender := models.Tender{
		Id:          primitive.NewObjectID(),
		Name:        tender.Name,
		Paket:       tender.Paket,
		Pagu:        tender.Pagu,
		Jadwal:      tender.Jadwal,
		Pelaksanaan: tender.Pelaksanaan,
		Pemilihan:   tender.Pemilihan,
		Pdn:         tender.Pdn,
		Idpagu:      tender.Idpagu,
	}

	result, err := tenderCollection.InsertOne(ctx, newTender)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllTender(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var tenders []models.Tender
	defer cancel()

	results, err := tenderCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singTender models.Tender
		if err = results.Decode(&singTender); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		tenders = append(tenders, singTender)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": tenders}},
	)
}

func GetFilterTender(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	var tenders []models.Tender
	defer cancel()
	//objId, _ := primitive.ObjectIDFromHex(paguId)
	// csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	results, err := tenderCollection.Find(ctx, bson.M{"idpagu": paguId})
	//err := paguCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&pagu)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleTender models.Tender
		if err = results.Decode(&singleTender); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		tenders = append(tenders, singleTender)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": tenders}},
	)
}

func DeleteTender(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	tenderId := c.Params("paguId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(tenderId)

	result, err := tenderCollection.DeleteOne(ctx, bson.M{"id": objId})
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

func EditTender(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	tenderId := c.Params("paguId")
	var tender models.Tender
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(tenderId)

	//validate the request body
	if err := c.BodyParser(&tender); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&tender); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": tender.Name, "paket": tender.Paket, "pagu": tender.Pagu, "jadwal": tender.Jadwal}

	result, err := tenderCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updateTender models.Tender
	if result.MatchedCount == 1 {
		err := tenderCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateTender)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updateTender}})
}

func GetTender(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	tenderId := c.Params("paguId")
	var tender models.Tender
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(tenderId)

	err := tenderCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&tender)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": tender}})
}

func GetTotalTender(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	var totalTenders []models.Totaltender
	defer cancel()

	matchStage := bson.D{{"$match", bson.D{{"idpagu", paguId}}}}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$name"},
			{"total", bson.D{{"$sum", 1}}},
			{"totalPagu", bson.D{{"$sum", "$pagu"}}},
		}},
	}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"name", "$_id"},
			{"total", 1},
			{"totalPagu", 1},
		}},
	}

	//totalPagu: {$sum: "$pagu"}
	/*
	   projectStage := bson.D{
	   	{"$project", bson.D{
	   	{"_id", 0},

	   	}},
	   	},
	   	}
	*/

	results, err := tenderCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleTender models.Totaltender
		if err = results.Decode(&singleTender); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		//fmt.Println(results)
		totalTenders = append(totalTenders, singleTender)
		//fmt.Print((tenders))
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": totalTenders}},
	)
}
