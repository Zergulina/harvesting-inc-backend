package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterEquipmentModelRoutes(app *fiber.App) {
	equipmentModel := app.Group("/api/equipments-types/:equipmentTypeId/equipments-models")

	equipmentModel.Get("", handler.GetMachineModels)
	equipmentModel.Post("", handler.CreateMachineModel)
	equipmentModel.Delete(":id", handler.DeleteMachineModel)
	equipmentModel.Put(":id", handler.UpdateMachineModel)
}
