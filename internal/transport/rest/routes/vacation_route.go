package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterVacationRoutes(app *fiber.App) {
	vacation := app.Group("/api/people:peopleId/vacations")

	vacation.Post("", handler.CreateVacation)
	vacation.Get("", handler.GetVacationsByPeopleId)
	vacation.Delete("", handler.DeleteVacation)
}
