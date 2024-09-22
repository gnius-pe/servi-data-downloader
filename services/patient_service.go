package services

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/gnius-pe/servi-data-downloader/models"
	"github.com/gnius-pe/servi-data-downloader/repositories"
)

func GetPatientById(id string) (*models.Patient, error) {
	return repositories.GetPatientByID(id)
}

func GetPatientAll() ([]models.Patient, error) {
	return repositories.GetAllPatients()
}

func GeneratePatientCSV(patients []models.Patient) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)
	headers := []string{"ID", "Name", "LastName", "NumberIdentification", "Email", "FirstNumberPhone", "SecondNumberPhone", "Sexo", "BirthDate", "Department", "Province", "District", "Reference", "AppointmentDate", "Specialties", "AppointmentDetail", "QuestionExamRecent", "SpiritualSupport", "FutureActivities", "Estate", "NumberFile", "CreatedAt", "UpdatedAt", "Version"}

	if err := writer.Write(headers); err != nil {
		return nil, err
	}

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
		if err := writer.Write(row); err != nil {
			return nil, err
		}
	}
	writer.Flush()
	return &buffer, nil
}
