package controllers

import (
	"github.com/gnius-pe/servi-data-downloader/services"
	"github.com/gnius-pe/servi-data-downloader/utils"
	"github.com/gofiber/fiber/v2"
)

func GetPatient(c *fiber.Ctx) error {
	id := c.Params("id")
	patient, err := services.GetPatientById(id) // Delegar a services
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}
	pdf, err := services.GeneratePatientPDF(patient) // Delegar la generación de PDF a services
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename="+patient.PersonalInformation.Name+"-"+patient.PersonalInformation.NumberIdentification+".pdf")
	return c.Send(pdf)
}

func ExportCSV(c *fiber.Ctx) error {
	patients, err := services.GetPatientAll()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	csvBuffer, err := services.GeneratePatientCSV(patients)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=patients.csv")
	return c.SendStream(csvBuffer)
}
