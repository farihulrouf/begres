package controllers

import (
	"begres/configs"
	"begres/models"
	"begres/responses"
	"context"

	//s"fmt"

	//"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var langsungCollection *mongo.Collection = configs.GetCollection(configs.DB, "langsung")
var validate = validator.New()

func CreateLangsung(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var langsung models.Langsung
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&langsung); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&langsung); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newLangsung := models.Langsung{
		Id:     primitive.NewObjectID(),
		Name:   langsung.Name,
		Paket:  langsung.Paket,
		Pagu:   langsung.Pagu,
		Jadwal: langsung.Jadwal,
		Pdn:    langsung.Pdn,
		Tipe:   langsung.Tipe,
		Ket:    langsung.Ket,
		Idpagu: langsung.Idpagu,
	}

	result, err := langsungCollection.InsertOne(ctx, newLangsung)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllLangsung(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var langsungs []models.Langsung
	defer cancel()

	results, err := langsungCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singLangsung models.Langsung
		if err = results.Decode(&singLangsung); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		langsungs = append(langsungs, singLangsung)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": langsungs}},
	)
}

func GetFilterLangsung(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	var langsungs []models.Langsung
	defer cancel()
	//objId, _ := primitive.ObjectIDFromHex(paguId)
	// csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	results, err := langsungCollection.Find(ctx, bson.M{"idpagu": paguId})
	//err := paguCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&pagu)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleLangsung models.Langsung
		if err = results.Decode(&singleLangsung); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		langsungs = append(langsungs, singleLangsung)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": langsungs}},
	)
}

func GetFilterLangsungByType(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	tipe := c.Params(("tipe"))
	var langsungs []models.Langsung
	defer cancel()
	//objId, _ := primitive.ObjectIDFromHex(paguId)
	// csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	//(ctx, User{Name: "UserName", Phone: "1234567890"})
	results, err := langsungCollection.Find(ctx, bson.M{"idpagu": paguId, "tipe": tipe})
	//err := paguCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&pagu)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleLangsung models.Langsung
		if err = results.Decode(&singleLangsung); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		langsungs = append(langsungs, singleLangsung)
	}

	return c.Status(http.StatusOK).JSON(
		responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": langsungs}},
	)
}

func DeleteLangsung(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	langsungId := c.Params("paguId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(langsungId)

	result, err := langsungCollection.DeleteOne(ctx, bson.M{"id": objId})
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

func EditLangsug(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	langsungId := c.Params("paguId")
	var langsung models.Langsung
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(langsungId)

	//validate the request body
	if err := c.BodyParser(&langsung); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&langsung); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	/*

		newLangsung := models.Langsung{
			Id:     primitive.NewObjectID(),
			Name:   langsung.Name,
			Paket:  langsung.Paket,
			Pagu:   langsung.Pagu,
			Jadwal: langsung.Jadwal,
			Pdn:    langsung.Pdn,
			Tipe:   langsung.Tipe,
			Ket:    langsung.Ket,
			Idpagu: langsung.Idpagu,
		}
	*/

	update := bson.M{"name": langsung.Name, "paket": langsung.Paket, "pagu": langsung.Pagu, "jadwal": langsung.Jadwal, "pdn": langsung.Pdn, "tipe": langsung.Tipe, "ket": langsung.Ket}

	result, err := langsungCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updateTender models.Tender
	if result.MatchedCount == 1 {
		err := langsungCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateTender)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updateTender}})
}

func GetLangsung(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	langsungId := c.Params("paguId")
	var langsung models.Langsung
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(langsungId)

	err := langsungCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&langsung)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": langsung}})
}

func GetTotalTenderLangsung(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	tipe := c.Params("tipe")
	var totalTenders []models.Totaltender
	defer cancel()
	//  { $sort : { total : 1 } }
	matchStage := bson.D{{"$match", bson.D{{"idpagu", paguId}}}}

	matchStageN := bson.D{{"$match", bson.D{{"tipe", tipe}}}}
	sortStage := bson.D{{"$sort", bson.D{{"total", 1}}}}
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

	// {$match: {"$and": [{idpagu: '635c8575a49fb441d6ff4670'}, {tipe: "langsung"}]}},

	results, err := langsungCollection.Aggregate(ctx, mongo.Pipeline{matchStage, matchStageN, groupStage, projectStage, sortStage})
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

func GetAllTotalTenderLangsung(c *fiber.Ctx) error {
	//fmt.Print("dieksekusi")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	paguId := c.Params("paguId")
	//tipe := c.Params("tipe")
	var totalTenders []models.Totaltipe
	defer cancel()
	matchStage := bson.D{{"$match", bson.D{{"idpagu", paguId}}}}
	sortStage := bson.D{{"$sort", bson.D{{"tipe", 1}}}}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$tipe"},
			{"total", bson.D{{"$sum", 1}}},
			{"totalPagu", bson.D{{"$sum", "$pagu"}}},
		}},
	}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"tipe", "$_id"},
			{"total", 1},
			{"totalPagu", 1},
		}},
	}

	results, err := langsungCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage, sortStage})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleTender models.Totaltipe
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

func GetAllTotalTenderLangsungAll(c *fiber.Ctx) error {
	//fmt.Print("dieksekusi")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//paguId := c.Params("paguId")
	//tipe := c.Params("tipe")
	var totalTenders []models.Totaltipe
	defer cancel()
	//matchStage := bson.D{{"$match", bson.D{{"idpagu", paguId}}}}
	sortStage := bson.D{{"$sort", bson.D{{"total", 1}}}}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$tipe"},
			{"total", bson.D{{"$sum", 1}}},
			{"totalPagu", bson.D{{"$sum", "$pagu"}}},
		}},
	}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"tipe", "$_id"},
			{"total", 1},
			{"totalPagu", 1},
		}},
	}

	results, err := langsungCollection.Aggregate(ctx, mongo.Pipeline{groupStage, projectStage, sortStage})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	//fmt.Print("disekskusi")
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleTender models.Totaltipe
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

func GetAllTotalTenderPdnAll(c *fiber.Ctx) error {
	//fmt.Print("dieksekusi")

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

	results, err := langsungCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage, sortStage})
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
