package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func GetPeople(c fiber.Ctx) error {
	people, err := repository.GetAllPeople(database.DB)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(people)
}

func DeletePeople(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsPeople(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	err = repository.DeletePeople(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdatePeople(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsPeople(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	cropType := new(models.People)
	if err := c.Bind().Body(cropType); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	cropType.Id = id
	cropType, err = repository.UpdatePeople(database.DB, cropType)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.JSON(cropType)
}
