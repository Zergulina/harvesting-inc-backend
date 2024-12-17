package routes

import (
	"backend/internal/transport/rest/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterPostRoutes(app *fiber.App) {
	crop := app.Group("/api/posts")

	crop.Get("", handler.GetPosts)
	crop.Post("", handler.CreatePost)
	crop.Delete(":id", handler.DeletePost)
	crop.Put(":id", handler.UpdatePost)
}
