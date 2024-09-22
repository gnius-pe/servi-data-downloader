package controllers

import (
	"bytes"
	"encoding/csv"
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

func GetPatient(c *fiber.Ctx) error {
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
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename="+patient.PersonalInformation.Name+"-"+patient.PersonalInformation.NumberIdentification+".pdf")
	return c.Send(pdf)
}

func ExportCSV(c *fiber.Ctx) error {
	patients, err := services.GetPatientAll()
	if err != nil {
		return utils.ErrorRespnse(c, fiber.StatusInternalServerError, err.Error())
	}
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)
	headers := []string{"ID", "Name", "LastName", "NumberIdentification", "Email", "FirstNumberPhone", "SecondNumberPhone", "Sexo", "BirthDate", "Department", "Province", "District", "Reference", "AppointmentDate", "Specialties", "AppointmentDetail", "QuestionExamRecent", "SpiritualSupport", "FutureActivities", "Estate", "NumberFile", "CreatedAt", "UpdatedAt", "Version"}
	writer.Write(headers)
	for _, patient := range patients {
		row := []string{
			patient.ID.Hex(),
			patient.PersonalInformation.Name,
			patient.PersonalInformation.LastName,
			patient.PersonalInformation.NumberIdentification,
			patient.PersonalInformation.Email,
			patient.PersonalInformation.FirstNumberPhone,
			patient.PersonalInformation.SecondNumberPhone,
			patient.PersonalInformation.Sexo,
			patient.PersonalInformation.BirthDate.Time().String(),
			patient.Location.Department,
			patient.Location.Province,
			patient.Location.District,
			patient.Location.Reference,
			patient.Cita.AppointmentDate.Time().String(),
			fmt.Sprintf("%v", patient.Cita.Specialties),
			patient.Cita.AppointmentDetail,
			fmt.Sprintf("%v", patient.Question.QuestionExamRecent),
			fmt.Sprintf("%v", patient.Question.SpiritualSupport),
			fmt.Sprintf("%v", patient.Question.FutureActivities),
			patient.Estate,
			fmt.Sprintf("%d", patient.NumberFile),
			patient.CreatedAt.Time().String(),
			patient.UpdatedAt.Time().String(),
			fmt.Sprintf("%d", patient.Version),
		}
		writer.Write(row)
	}
	writer.Flush()
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=patients.csv")
	return c.SendStream(&buffer)
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
	birthTime := time.Unix(int64(birthDate)/1000, 0)
	currentTime := time.Now()
	age := currentTime.Year() - birthTime.Year()
	if currentTime.YearDay() < birthTime.YearDay() {
		age--
	}
	return age
}

// Función para parsear y formatear una fecha
func parseDate(date primitive.DateTime) string {
	timeDate := time.Unix(int64(date)/1000, 0)
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
