package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterCustomerRoutes(app *fiber.App) {
	cropType := app.Group("/api/customers")

	cropType.Get("", handler.GetCustomers)
	cropType.Post("", handler.CreateCustomer)
	cropType.Delete(":id", handler.DeleteCustomer)
	cropType.Put(":id", handler.UpdateCustomer)
	cropType.Patch(":id", handler.PatchCustomer)
}
