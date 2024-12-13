package handler

import (
	"backend/internal/database"

	"github.com/gofiber/fiber/v2"
)

func ResetHandler(c *fiber.Ctx) error {
	database.ResetDb()
	return c.SendString("База данных успешно сброшена")
}
