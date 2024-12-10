package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v3"
)

func RegisterCropTypeRoutes(app *fiber.App) {
	crop := app.Group("/api/crop-types/:cropTypeId/crops")

	crop.Get("", handler.GetCropTypes)
	crop.Post("", handler.CreateCropType)
	crop.Delete(":id", handler.DeleteCropType)
	crop.Put(":id", handler.UpdateCropType)
}
