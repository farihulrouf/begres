package routes

import (
	"begres/controllers"
	"begres/middleware"

	"github.com/gofiber/fiber/v2"
)

func AppRoute(app *fiber.App) {

	app.Post("/api/users/singup", controllers.SingUp)
	app.Post("/api/users/login", controllers.Login)
	app.Get("/api/users/getall", middleware.Authentication, controllers.GetAllUser)
	app.Delete("/api/user/:paguId", middleware.Authentication, controllers.DeleteUser)
	app.Get("/api/user/:paguId", middleware.Authentication, controllers.GetUser)
	app.Put("/api/user/:paguId", middleware.Authentication, controllers.Edituser)

	app.Post("/api/link", middleware.Authentication, controllers.CreateLink)
	app.Get("/api/link", middleware.Authentication, controllers.GetAllLink)
	app.Put("/api/link/:paguId", middleware.Authentication, controllers.EditLink)
	app.Get("/api/link/:paguId", middleware.Authentication, controllers.GetLink)
	app.Delete("/api/link/:paguId", middleware.Authentication, controllers.DeleteLink)

	//Post Pagu
	app.Post("/api/pagus", middleware.Authentication, controllers.CreatePagu)
	app.Get("/api/pagus", middleware.Authentication, controllers.GetAllPagu)
	app.Get("/api/pagus/filter", middleware.Authentication, controllers.GetAllFilter)
	app.Get("/api/pagus/:paguId", middleware.Authentication, controllers.GetPagu)
	app.Delete("/api/pagus/:paguId", middleware.Authentication, controllers.DeletePagu)
	app.Put("/api/pagus/:paguId", middleware.Authentication, controllers.EditPagu)
	app.Put("/api/pagus/edit/:paguId", middleware.Authentication, controllers.EditPaguUpload)
	//EditPaguUpload

	//Pagu anggaran
	app.Post("/api/anggaran", middleware.Authentication, controllers.CreateAnggaran)
	app.Get("/api/anggaran", middleware.Authentication, controllers.GetAllAnggaran)
	app.Get("/api/anggaran/pagu/:paguId", middleware.Authentication, controllers.GetFilterAnggaran)
	app.Delete("/api/anggaran/:paguId", middleware.Authentication, controllers.DeleteAnggran)
	app.Put("/api/anggaran/:paguId", middleware.Authentication, controllers.EditAnggaran)
	app.Get("/api/anggaran/:paguId", middleware.Authentication, controllers.GetAnggaran)
	app.Get("/api/anggaran/pagu/total/:paguId", middleware.Authentication, controllers.GetAllTotalPaguAnggaran)
	//GetAllTotalPaguAnggaran

	//langsung
	app.Post("/api/langsung", middleware.Authentication, controllers.CreateLangsung)
	app.Get("/api/langsung", middleware.Authentication, controllers.GetAllLangsung)
	app.Get("/api/langsung/pagu/:paguId", middleware.Authentication, controllers.GetFilterLangsung)
	app.Get("/api/langsung/pagu/:paguId/:tipe", middleware.Authentication, controllers.GetFilterLangsungByType)
	app.Get("/api/langsung/:paguId", middleware.Authentication, controllers.GetLangsung)
	app.Get("/api/langsung/pagu/total/:paguId/:tipe", middleware.Authentication, controllers.GetTotalTenderLangsung)
	app.Get("/api/langsung/totalsemua/:paguId", middleware.Authentication, controllers.GetAllTotalTenderLangsung)

	app.Get("/api/langsung/totalseleksitender/:paguId", middleware.Authentication, controllers.GetAllTotalTenderLangsungBySeleksiCepat)
	//GetAllTotalTenderCepatAllSeleksi
	app.Get("/api/jumlahtotalseleksi", middleware.Authentication, controllers.GetAllTotalTenderCepatAllSeleksi)
	app.Get("/api/jumlahtotal", middleware.Authentication, controllers.GetAllTotalTenderLangsungAll)
	app.Get("/api/totalpdn/:paguId", middleware.Authentication, controllers.GetAllTotalTenderPdnAll)

	//app.Get("/api/jumlahratapdn/:paguId", controllers.GetAllTotalTenderPdnLookup)
	//GetAllTotalTenderPdnLookup
	app.Delete("/api/langsung/:paguId", middleware.Authentication, controllers.DeleteLangsung)
	app.Put("/api/langsung/:paguId", middleware.Authentication, controllers.EditLangsug)

	//uploudfile
	app.Post("/api/uploud", controllers.CreatetoUpload)
	app.Get("/api/upload/:paguId", controllers.GetFilterUpload)

}
