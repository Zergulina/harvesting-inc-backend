package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterMachineEquipmentTypeRoutes(app *fiber.App) {
	machineType := app.Group("/api/machine-types/:machineTypeId/equipments-types")

	machineType.Get("", handler.GetAllEquipmentTypesByMachineTypeId)
	machineType.Post(":equipmentTypeId", handler.CreateMachineEquipmentType)
	machineType.Delete(":equipmentTypeId", handler.DeleteMachineEquipmentType)
}
