/*
package main

import (
	"github.com/gnius-pe/servi-data-downloader/configs"
	"github.com/gnius-pe/servi-data-downloader/controllers"
	"github.com/gnius-pe/servi-data-downloader/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	utils.ConnectDB()
	app := fiber.New()
	app.Use(cors.New())
	controllers.SetupRoutes(app)

	app.Listen(configs.ServerPort)
}

*/

package main

import (
	"context"
	"fmt"

	"github.com/gnius-pe/servi-data-downloader/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://doadmin:607z2e18cw3BEKk4@db-dar-d8c80c07.mongo.ondigitalocean.com"))
	if err != nil {
		panic((err))
	}

	collectionPatient := client.Database("test").Collection("testpatients")

	app.Use(cors.New())

	app.Get("/api/patient/downloader/:id", func(c *fiber.Ctx) error {
		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "ID invalido",
			})
		}

		var patient models.Patient
		err = collectionPatient.FindOne(c.Context(), bson.M{"_id": id}).Decode(&patient)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON((fiber.Map{
				"error": "Paciente no encontrado",
			}))
		}
		return c.JSON(patient)
	})
	app.Listen(":3010")
	fmt.Println("Server on port 3000")
}
