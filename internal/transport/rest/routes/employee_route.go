package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterEmployeeRoutes(app *fiber.App) {
	cropType := app.Group("/api/people/:peopleId/employees")

	cropType.Get("", handler.GetEmployees)
	cropType.Post("", handler.CreateEmployee)
	cropType.Delete(":postId", handler.DeleteEmployee)
	cropType.Put(":postId", handler.UpdateEmployee)
}
