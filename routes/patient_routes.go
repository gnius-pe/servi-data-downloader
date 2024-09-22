package routes

import (
	"github.com/gnius-pe/servi-data-downloader/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetupPatientRoutes define todas las rutas relacionadas con los pacientes.
func SetupPatientRoutes(app *fiber.App) {
	app.Get("/api/patient/downloader/:id", controllers.GetPatient)
	app.Get("/api/patient/download/csv", controllers.ExportCSV)
}
