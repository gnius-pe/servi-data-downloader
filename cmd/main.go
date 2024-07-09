package main

import (
	"fmt"
	"os"

	"github.com/gnius-pe/servi-data-downloader/configs"
	"github.com/gnius-pe/servi-data-downloader/controllers"
	"github.com/gnius-pe/servi-data-downloader/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	configs.LoadConfig()

	fmt.Println("PATH:", os.Getenv("PATH"))
	err := utils.ConnectDB()
	if err != nil {
		panic(err) // o manejar el error de alguna otra manera
	}
	app := fiber.New()
	app.Use(cors.New())
	controllers.SetupRoutes(app)
	app.Listen(configs.ServerPort)
}
