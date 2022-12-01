package routes

import (
	"begres/controllers"

	"github.com/gofiber/fiber/v2"
)

func AppRoute(app *fiber.App) {

	app.Post("/api/users/singup", controllers.SingUp)

	//Post Pagu
	app.Post("/api/pagus", controllers.CreatePagu)
	app.Get("/api/pagus", controllers.GetAllPagu)
	app.Get("/api/pagus/filter", controllers.GetAllFilter)
	app.Get("/api/pagus/:paguId", controllers.GetPagu)
	app.Delete("/api/pagus/:paguId", controllers.DeletePagu)
	app.Put("/api/pagus/:paguId", controllers.EditPagu)
	app.Put("/api/pagus/edit/:paguId", controllers.EditPaguUpload)
	//EditPaguUpload

	//Pagu anggaran
	app.Post("/api/anggaran", controllers.CreateAnggaran)
	app.Get("/api/anggaran", controllers.GetAllAnggaran)
	app.Get("/api/anggaran/pagu/:paguId", controllers.GetFilterAnggaran)
	app.Delete("/api/anggaran/:paguId", controllers.DeleteAnggran)
	app.Put("/api/anggaran/:paguId", controllers.EditAnggaran)
	app.Get("/api/anggaran/:paguId", controllers.GetAnggaran)

	//langsung
	app.Post("/api/langsung", controllers.CreateLangsung)
	app.Get("/api/langsung", controllers.GetAllLangsung)
	app.Get("/api/langsung/pagu/:paguId", controllers.GetFilterLangsung)
	app.Get("/api/langsung/pagu/:paguId/:tipe", controllers.GetFilterLangsungByType)
	app.Get("/api/langsung/:paguId", controllers.GetLangsung)
	app.Get("/api/langsung/pagu/total/:paguId/:tipe", controllers.GetTotalTenderLangsung)
	app.Get("/api/langsung/totalsemua/:paguId", controllers.GetAllTotalTenderLangsung)

	app.Get("/api/langsung/totalseleksitender/:paguId", controllers.GetAllTotalTenderLangsungBySeleksiCepat)
	//GetAllTotalTenderCepatAllSeleksi
	app.Get("/api/jumlahtotalseleksi", controllers.GetAllTotalTenderCepatAllSeleksi)
	app.Get("/api/jumlahtotal", controllers.GetAllTotalTenderLangsungAll)
	app.Get("/api/totalpdn/:paguId", controllers.GetAllTotalTenderPdnAll)

	//app.Get("/api/jumlahratapdn/:paguId", controllers.GetAllTotalTenderPdnLookup)
	//GetAllTotalTenderPdnLookup
	app.Delete("/api/langsung/:paguId", controllers.DeleteLangsung)
	app.Put("/api/langsung/:paguId", controllers.EditLangsug)

	//uploudfile
	app.Post("/api/uploud", controllers.CreatetoUpload)
	app.Get("/api/upload/:paguId", controllers.GetFilterUpload)

}
