package rest

import (
	"backend/internal/transport/rest/routes"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App) {
	routes.RegisterCropTypeRoutes(app)
	routes.RegisterCropRoutes(app)
	routes.RegisterPeopleRoutes(app)
	routes.RegisterAccountRoutes(app)
}
