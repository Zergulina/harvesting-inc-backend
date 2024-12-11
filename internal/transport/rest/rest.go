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
}
