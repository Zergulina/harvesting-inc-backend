package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterMachineModelRoutes(app *fiber.App) {
	machineModel := app.Group("/api/machine-types/:machineTypeId/machine-models")

	machineModel.Get("", handler.GetMachineModels)
	machineModel.Post("", handler.CreateMachineModel)
	machineModel.Delete(":id", handler.DeleteMachineModel)
	machineModel.Put(":id", handler.UpdateMachineModel)
}
