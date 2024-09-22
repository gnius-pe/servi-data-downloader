package services

import (
	"bytes"
	"text/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gnius-pe/servi-data-downloader/models"
	"github.com/gnius-pe/servi-data-downloader/utils"
)

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

func renderHTML(data interface{}) (string, error) {
	tmpl := template.Must(template.New("patient.html").Funcs(template.FuncMap{
		"concat":                       utils.Concat,
		"getCurrentDateTime":           utils.GetCurrentDateTime,
		"calculateAge":                 utils.CalculateAge,
		"parseDate":                    utils.ParseDate,
		"booleanToString":              utils.BooleanToString,
		"concatenateSpecialtiesLabels": utils.ConcatenateSpecialtiesLabels,
	}).ParseFiles("./templates/patient.html"))

	var html bytes.Buffer
	if err := tmpl.Execute(&html, data); err != nil {
		return "", err
	}
	return html.String(), nil
}

func GeneratePatientPDF(patient *models.Patient) ([]byte, error) {
	html, err := renderHTML(patient)
	if err != nil {
		return nil, err
	}

	pdf, err := generatePDF(html)
	if err != nil {
		return nil, err
	}
	return pdf, nil
}
