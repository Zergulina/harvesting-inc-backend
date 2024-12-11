package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterFieldTypeRoutes(app *fiber.App) {
	crop := app.Group("/api/customers/:customerId/fields")

	crop.Get("", handler.GetFields)
	crop.Post("", handler.CreateField)
	crop.Delete(":id", handler.DeleteField)
	crop.Put(":id", handler.UpdateField)
}
