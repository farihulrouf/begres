package main

import (
	"begres/configs"
	"begres/routes" //add this

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	//run database

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))
	app.Static("/", "./public")

	configs.ConnectDB()

	//routes
	routes.AppRoute(app) //add this

	app.Listen(":3000")

}
