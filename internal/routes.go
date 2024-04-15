package internal

import (
	"github.com/HamzahAA15/records_api/internal/records"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// setup from muliple routes of domain
func SetupRoutes(app *fiber.App, db *sqlx.DB) {
	recordsRepository := records.NewRecordsRepository(db)
	recordsService := records.NewRecordsService(recordsRepository)
	recordsHandler := records.NewRecordsHandler(recordsService)
	recordsHandler.SetupRoutes(app)
}
