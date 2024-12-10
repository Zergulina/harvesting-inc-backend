package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v3"
)

func RegisterCropRoutes(app *fiber.App) {
	cropType := app.Group("/api/crop-types")

	cropType.Get("", handler.GetCropTypes)
	cropType.Post("", handler.CreateCropType)
	cropType.Delete(":id", handler.DeleteCropType)
	cropType.Put(":id", handler.UpdateCropType)
}
