package services

import (
	"github.com/gnius-pe/servi-data-downloader/models"
	"github.com/gnius-pe/servi-data-downloader/repositories"
)

func GetPatientById(id string) (*models.Patient, error) {
	return repositories.GetPatientByID(id)
}

func GetPatientAll() ([]models.Patient, error) {
	return repositories.GetAllPatients()
}
