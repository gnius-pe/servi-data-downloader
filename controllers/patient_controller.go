package controllers

import (
	"github.com/gnius-pe/servi-data-downloader/services"
	"github.com/gnius-pe/servi-data-downloader/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/patient/downloader/:id", getPatient)
}

func getPatient(c *fiber.Ctx) error {
	id := c.Params("id")
	patient, err := services.GetPatientById(id)
	if err != nil {
		return utils.ErrorRespnse(c, fiber.StatusNotFound, err.Error())
	}
	return c.JSON(patient)
}
