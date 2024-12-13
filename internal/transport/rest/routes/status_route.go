package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterStatusRoutes(app *fiber.App) {
	cropType := app.Group("/api/statuses")

	cropType.Get("", handler.GetStatuses)
	cropType.Post("", handler.CreateStatus)
	cropType.Delete(":id", handler.DeleteStatus)
	cropType.Put(":id", handler.UpdateStatus)
}
