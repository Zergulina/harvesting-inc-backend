package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterAccountRoutes(app *fiber.App) {
	account := app.Group("/api/account")

	account.Post("register", handler.Register)
	account.Post("login", handler.Login)
}
