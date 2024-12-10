package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func GetCropTypes(c fiber.Ctx) error {
	cropTypes, err := repository.GetAllCropTypes(database.DB)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(cropTypes)
}

func CreateCropType(c fiber.Ctx) error {
	cropType := new(models.CropType)
	if err := c.Bind().Body(cropType); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	cropType, err := repository.CreateCropType(database.DB, cropType)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(cropType)
}

func DeleteCropType(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsCropType(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	err = repository.DeleteCropType(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateCropType(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsCropType(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	cropType := new(models.CropType)
	if err := c.Bind().Body(cropType); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	cropType.Id = id
	cropType, err = repository.UpdateCropType(database.DB, cropType)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.JSON(cropType)
}
