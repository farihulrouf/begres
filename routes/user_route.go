package routes

import (
	"begres/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Put("/user/:userId", controllers.EditAUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)

	//Post Pagu

	app.Post("/api/pagus", controllers.CreatePagu)
	app.Get("/api/pagus", controllers.GetAllPagu)
	app.Get("/api/pagus/:paguId", controllers.GetPagu)
	app.Delete("/api/pagus/:paguId", controllers.DeletePagu)
	app.Put("/api/pagus/:paguId", controllers.EditPagu)

	//Pagu anggaran
	app.Post("/api/anggaran", controllers.CreatePanggaran)
	app.Get("/api/anggaran", controllers.GetAllAnggaran)

	//tender
	app.Post("/api/tender", controllers.CreateTender)
	app.Get("/api/tender", controllers.GetAllTender)

	//langsung
	app.Post("/api/langsung", controllers.CreateLangsung)
	app.Get("/api/langsung", controllers.GetAllLangsung)

}
