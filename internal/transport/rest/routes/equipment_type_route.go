package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterEquipmentTypeRoutes(app *fiber.App) {
	equipmentType := app.Group("/api/equipments-types")

	equipmentType.Get("", handler.GetMachineTypes)
	equipmentType.Post("", handler.CreateMachineType)
	equipmentType.Delete(":id", handler.DeleteMachineType)
	equipmentType.Put(":id", handler.UpdateMachineType)
}
