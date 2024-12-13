package rest

import (
	"backend/internal/transport/rest/routes"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	routes.RegisterAccountRoutes(app)
}

func RegisterProtectedRoutes(app *fiber.App) {
	routes.RegisterCropTypeRoutes(app)
	routes.RegisterCropRoutes(app)
	routes.RegisterPeopleRoutes(app)
	routes.RegisterEmployeeRoutes(app)
	routes.RegisterCustomerRoutes(app)
	routes.RegisterFieldTypeRoutes(app)
	routes.RegisterStatusRoutes(app)
	routes.RegisterMachineTypeRoutes(app)
	routes.RegisterMachineModelRoutes(app)
	routes.RegisterMachineRoutes(app)
	routes.RegisterEquipmentTypeRoutes(app)
	routes.RegisterEquipmentModelRoutes(app)
	routes.RegisterEquipmentRoutes(app)
	routes.RegisterMachineEquipmentTypeRoutes(app)
	routes.RegisterWorkRoutes(app)
	routes.RegisterWorkTripRoutes(app)
	routes.RegisterVacationRoutes(app)
	routes.RegisterAdminRoutes(app)
}
