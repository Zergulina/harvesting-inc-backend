package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterPeopleRoutes(app *fiber.App) {
	people := app.Group("/api/people")

	people.Get("", handler.GetPeople)
	people.Delete(":id", handler.DeletePeople)
	people.Put(":id", handler.UpdatePeople)
	people.Get("me", handler.Me)
}
