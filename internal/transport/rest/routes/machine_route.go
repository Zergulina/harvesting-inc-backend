package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterMachineRoutes(app *fiber.App) {
	machine := app.Group("/api/machine-types/:machineTypeId/machine-models/:machineModelId/machines")

	machine.Get("", handler.GetMachines)
	machine.Post("", handler.CreateMachine)
	machine.Delete(":invNumber", handler.DeleteMachine)
	machine.Put(":invNumber", handler.UpdateMachine)
}
