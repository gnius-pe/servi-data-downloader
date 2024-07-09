package controllers

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
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

	html, err := renderHTML(patient)

	if err != nil {
		return utils.ErrorRespnse(c, fiber.StatusInternalServerError, err.Error())
	}

	pdf, err := generatePDF(html)

	if err != nil {
		return utils.ErrorRespnse(c, fiber.StatusInternalServerError, err.Error())
	}
	fmt.Println("estapa de envio", pdf)
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=patient.pdf")
	return c.Send(pdf)
}

func renderHTML(data interface{}) (string, error) {
	tmpl, err := template.ParseFiles("./templates/patient.html")
	if err != nil {
		return "", err
	}

	var html bytes.Buffer
	if err := tmpl.Execute(&html, data); err != nil {
		return "", err
	}
	return html.String(), nil
}

func generatePDF(html string) ([]byte, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(html))))

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
