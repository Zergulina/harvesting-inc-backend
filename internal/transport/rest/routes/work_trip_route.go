package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterWorkTripRoutes(app *fiber.App) {
	workTrip := app.Group("/api/customers/:customerId/fields/:fieldId/works/:workId/workTrips")

	workTrip.Post("", handler.DeleteWorkTrip)
	workTrip.Put(":id", handler.UpdateWorkTrip)
	workTrip.Delete(":id", handler.DeleteWorkTrip)

	app.Get("/api/works/:workId/workTrips", handler.GetWorkTripsByWorkId)
}
