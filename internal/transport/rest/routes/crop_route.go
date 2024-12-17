package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterCropRoutes(app *fiber.App) {
	crop := app.Group("/api/crop-types/:cropTypeId/crops")

	crop.Get("", handler.GetCrops)
	crop.Post("", handler.CreateCrop)
	crop.Delete(":id", handler.DeleteCrop)
	crop.Put(":id", handler.UpdateCrop)
}
