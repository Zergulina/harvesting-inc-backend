package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterEquipmentRoutes(app *fiber.App) {
	equipment := app.Group("/api/equipment-types/:equipmentTypeId/equipment-models/:equipmentModelId/equipments")

	equipment.Get("", handler.GetMachines)
	equipment.Post("", handler.CreateMachine)
	equipment.Delete(":invNumber", handler.DeleteMachine)
	equipment.Put(":invNumber", handler.UpdateMachine)
}
