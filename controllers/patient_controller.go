package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gnius-pe/servi-data-downloader/models"
	"github.com/gnius-pe/servi-data-downloader/services"
	"github.com/gnius-pe/servi-data-downloader/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fmt.Println(patient.PersonalInformation.BirthDate)

	html, err := renderHTML(patient)

	if err != nil {
		return utils.ErrorRespnse(c, fiber.StatusInternalServerError, err.Error())
	}

	pdf, err := generatePDF(html)

	if err != nil {
		return utils.ErrorRespnse(c, fiber.StatusInternalServerError, err.Error())
	}
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=patient.pdf")
	return c.Send(pdf)
}

// funciones
// Función auxiliar para concatenar strings
func concat(args ...string) string {
	return strings.Join(args, "")
}

// Función auxiliar para obtener la fecha y hora actual en formato "dia mes año y hora"
func getCurrentDateTime() string {
	return time.Now().Format("02-01-2006 15:04")
}

// Función para calcular la edad a partir de la fecha de nacimiento
func calculateAge(birthDate primitive.DateTime) int {
	// Convertir primitive.DateTime a time.Time
	birthTime := time.Unix(int64(birthDate)/1000, 0) // Convertir milisegundos a segundos

	// Obtener la fecha actual
	currentTime := time.Now()

	// Calcular la diferencia de años
	age := currentTime.Year() - birthTime.Year()

	// Ajustar la edad si aún no se ha cumplido el cumpleaños este año
	if currentTime.YearDay() < birthTime.YearDay() {
		age--
	}

	return age
}

// Función para parsear y formatear una fecha
func parseDate(date primitive.DateTime) string {
	// Convertir primitive.DateTime a time.Time
	timeDate := time.Unix(int64(date)/1000, 0) // Convertir milisegundos a segundos

	// Formatear la fecha como mes día año
	formattedDate := timeDate.Format("07/02/2006")

	return formattedDate
}

func booleanToString(value bool) string {
	if value {
		return "Si"
	}
	return "No"
}

func concatenateSpecialtiesLabels(specialties []models.Specialty) string {
	var labels []string
	for _, specialty := range specialties {
		labels = append(labels, specialty.Label)
	}
	return strings.Join(labels, " - ")
}

func renderHTML(data interface{}) (string, error) {
	/**
	tmpl, err := template.ParseFiles("./templates/patient.html")
	if err != nil {
		return "", err
	}
	*/

	tmpl := template.Must(template.New("patient.html").Funcs(template.FuncMap{
		"concat":                       concat,
		"getCurrentDateTime":           getCurrentDateTime,
		"calculateAge":                 calculateAge,
		"parseDate":                    parseDate,
		"booleanToString":              booleanToString,
		"concatenateSpecialtiesLabels": concatenateSpecialtiesLabels,
	}).ParseFiles("./templates/patient.html"))

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

	page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(html)))
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
