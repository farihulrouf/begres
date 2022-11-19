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

var anggaranCollection *mongo.Collection = configs.GetCollection(configs.DB, "anggaran")

func CreateAnggaran(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var panggaran models.Anggaran
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&panggaran); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&panggaran); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newPanggaran := models.Anggaran{
		Id:     primitive.NewObjectID(),
		Name:   panggaran.Name,
		Pagu:   panggaran.Pagu,
		Paket:  panggaran.Paket,
		Jadwal: panggaran.Jadwal,
		Pdn:    panggaran.Pdn,
		Idpagu: panggaran.Idpagu,
	}

	result, err := anggaranCollection.InsertOne(ctx, newPanggaran)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllAnggaran(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var panggarans []models.Anggaran
	defer cancel()

	results, err := anggaranCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singPanggaran models.Anggaran
		if err = results.Decode(&singPanggaran); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		panggarans = append(panggarans, singPanggaran)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": panggarans}},
	)
}
func GetFilterAnggaran(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	var panggarans []models.Anggaran
	defer cancel()
	sortStage := bson.D{{"$sort", bson.D{{"name", 1}}}}
	matchStage := bson.D{{"$match", bson.D{{"idpagu", paguId}}}}
	//results, err := anggaranCollection.Find(ctx, bson.M{"idpagu": paguId})

	results, err := anggaranCollection.Aggregate(ctx, mongo.Pipeline{matchStage, sortStage})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singPanggaran models.Anggaran
		if err = results.Decode(&singPanggaran); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		panggarans = append(panggarans, singPanggaran)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": panggarans}},
	)
}

func DeleteAnggran(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	anggaranId := c.Params("paguId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(anggaranId)

	result, err := anggaranCollection.DeleteOne(ctx, bson.M{"id": objId})
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

func EditAnggaran(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	anggaranId := c.Params("paguId")
	var anggaran models.Anggaran
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(anggaranId)

	//validate the request body
	if err := c.BodyParser(&anggaran); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&anggaran); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": anggaran.Name, "paket": anggaran.Paket, "pagu": anggaran.Pagu, "jadwal": anggaran.Jadwal, "pdn": anggaran.Pdn, "idpagu": anggaran.Idpagu}

	result, err := anggaranCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updateAnggaran models.Anggaran
	if result.MatchedCount == 1 {
		err := anggaranCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateAnggaran)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updateAnggaran}})
}

func GetAnggaran(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	anggaranId := c.Params("paguId")
	var tender models.Anggaran
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(anggaranId)

	err := anggaranCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&tender)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": tender}})
}

func GetAllTotalTenderPdnAllTender(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	//tipe := c.Params("tipe")
	var totalTenders []models.Totaltender
	defer cancel()
	matchStage := bson.D{{"$match", bson.D{{"idpagu", paguId}}}}
	sortStage := bson.D{{"$sort", bson.D{{"name", 1}}}}
	//{$divide: ["$details.salary", 2]}}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$name"},
			{"total", bson.D{{"$sum", 1}}},
			{"totalPagu", bson.D{{"$sum", "$pagu"}}},
			{"pdn", bson.D{{"$avg", "$pdn"}}},
		}},
	}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"name", "$_id"},
			{"total", 1},
			{"totalPagu", 1},
			{"pdn", 1},
			{"idpagu", 1},
		}},
	}

	results, err := anggaranCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage, sortStage})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleTender models.Totaltender
		if err = results.Decode(&singleTender); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		//fmt.Println(results)
		totalTenders = append(totalTenders, singleTender)
		//fmt.Print((totalTenders))
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": totalTenders}},
	)
}
