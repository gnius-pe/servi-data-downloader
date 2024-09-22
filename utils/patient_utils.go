package utils

import (
	"strings"
	"time"

	"github.com/gnius-pe/servi-data-downloader/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CalculateAge(birthDate primitive.DateTime) int {
	birthTime := time.Unix(int64(birthDate)/1000, 0)
	currentTime := time.Now()
	age := currentTime.Year() - birthTime.Year()
	if currentTime.YearDay() < birthTime.YearDay() {
		age--
	}
	return age
}

func ParseDate(date primitive.DateTime) string {
	timeDate := time.Unix(int64(date)/1000, 0)
	return timeDate.Format("07/02/2006")
}

func BooleanToString(value bool) string {
	if value {
		return "Si"
	}
	return "No"
}

func ConcatenateSpecialtiesLabels(specialties []models.Specialty) string {
	var labels []string
	for _, specialty := range specialties {
		labels = append(labels, specialty.Label)
	}
	return strings.Join(labels, " - ")
}
