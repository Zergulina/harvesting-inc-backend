package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterReportRoutes(app *fiber.App) {
	reports := app.Group("api/reports/")

	reports.Get("field-harvesting", handler.GetFieldHarvestingReport)
	reports.Get("people-experience", handler.GetPeopleExperienceReport)
	reports.Get("machines", handler.GetMachinesReport)
}
