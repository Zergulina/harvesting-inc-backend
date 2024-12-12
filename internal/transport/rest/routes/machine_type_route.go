package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterMachineTypeRoutes(app *fiber.App) {
	machineType := app.Group("/api/machine-types")

	machineType.Get("", handler.GetMachineTypes)
	machineType.Post("", handler.CreateMachineType)
	machineType.Delete(":id", handler.DeleteMachineType)
	machineType.Put(":id", handler.UpdateMachineType)
}
