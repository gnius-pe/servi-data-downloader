package main

import (
	"github.com/gnius-pe/servi-data-downloader/configs"
	"github.com/gnius-pe/servi-data-downloader/routes"
	"github.com/gnius-pe/servi-data-downloader/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	configs.LoadConfig()
	err := utils.ConnectDB()
	if err != nil {
		panic(err) // o manejar el error de alguna otra manera
	}
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	routes.SetupPatientRoutes(app)
	app.Listen(configs.ServerPort)
}
