package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterWorkRoutes(app *fiber.App) {
	work := app.Group("/api/customers/:customerId/fields/:fieldId/works")

	work.Post("", handler.DeleteWork)
	work.Put(":id", handler.UpdateWork)
	work.Delete(":id", handler.DeleteWork)

	app.Get("/api/customers/:customerId/works", handler.GetWorksByCustomerId)
}
