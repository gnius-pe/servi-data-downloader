package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/api/patient/downloader", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "patients desde el backend",
		})
	})
	app.Listen(":3010")
	fmt.Println("Server on port 3000")
}
