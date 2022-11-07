package routes

import (
	"begres/controllers"

	"github.com/gofiber/fiber/v2"
)

func AppRoute(app *fiber.App) {
	//Post Pagu

	app.Post("/api/pagus", controllers.CreatePagu)
	app.Get("/api/pagus", controllers.GetAllPagu)
	app.Get("/api/pagus/:paguId", controllers.GetPagu)
	app.Delete("/api/pagus/:paguId", controllers.DeletePagu)
	app.Put("/api/pagus/:paguId", controllers.EditPagu)

	//Pagu anggaran
	app.Post("/api/anggaran", controllers.CreatePanggaran)
	app.Get("/api/anggaran", controllers.GetAllAnggaran)
	app.Get("/api/anggaran/pagu/:paguId", controllers.GetFilterAnggaran)
	app.Delete("/api/anggaran/:paguId", controllers.DeleteAnggran)
	app.Put("/api/anggaran/:paguId", controllers.EditAnggaran)
	app.Get("/api/anggaran/:paguId", controllers.GetAnggaran)

	//tender
	app.Post("/api/tender", controllers.CreateTender)
	app.Get("/api/tender", controllers.GetAllTender)
	app.Put("/api/tender/:paguId", controllers.EditTender)
	app.Delete("/api/tender/:paguId", controllers.DeleteTender)
	app.Get("/api/tender/pagu/:paguId", controllers.GetFilterTender)
	app.Get("/api/tender/:paguId", controllers.GetTender)
	app.Get("/api/tender/total/:paguId", controllers.GetTotalTender)
	app.Get("/api/tender/totalpaket/:paguId", controllers.GetTotalTenderName)
	app.Get("/api/jumlahtender", controllers.GetTotalTenderNameAll)

	//langsung
	app.Post("/api/langsung", controllers.CreateLangsung)
	app.Get("/api/langsung", controllers.GetAllLangsung)
	app.Get("/api/langsung/pagu/:paguId", controllers.GetFilterLangsung)
	app.Get("/api/langsung/pagu/:paguId/:tipe", controllers.GetFilterLangsungByType)
	app.Get("/api/langsung/:paguId", controllers.GetLangsung)
	app.Get("/api/langsung/pagu/total/:paguId/:tipe", controllers.GetTotalTenderLangsung)
	app.Get("/api/langsung/totalsemua/:paguId", controllers.GetAllTotalTenderLangsung)
	app.Get("/api/jumlahtotal", controllers.GetAllTotalTenderLangsungAll)

	app.Delete("/api/langsung/:paguId", controllers.DeleteLangsung)
	app.Put("/api/kecuali/:paguId", controllers.EditLangsug)

	//

}
