package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v3"
)

func RegisterPeopleRoutes(app *fiber.App) {
	crop := app.Group("/api/people")

	crop.Get("", handler.GetPeople)
	crop.Delete(":id", handler.DeletePeople)
	crop.Put(":id", handler.UpdatePeople)
}
