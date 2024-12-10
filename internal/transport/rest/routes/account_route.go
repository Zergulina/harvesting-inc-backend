package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v3"
)

func RegisterAccountRoutes(app *fiber.App) {
	crop := app.Group("/api/account")

	crop.Post("register", handler.Register)
	crop.Post("login", handler.Login)
}
