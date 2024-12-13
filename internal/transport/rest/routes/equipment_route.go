package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterEquipmentRoutes(app *fiber.App) {
	equipments := app.Group("/api/equipments-types/:equipmentTypeId/equipments-models/:equipmentModelId/equipments")

	equipments.Get("", handler.GetMachines)
	equipments.Post("", handler.CreateMachine)
	equipments.Delete(":invNumber", handler.DeleteMachine)
	equipments.Put(":invNumber", handler.UpdateMachine)
}
