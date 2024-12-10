package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func GetCrops(c fiber.Ctx) error {
	cropTypes, err := repository.GetAllCrops(database.DB)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(cropTypes)
}

func CreateCrop(c fiber.Ctx) error {
	cropTypeId, err := strconv.ParseUint(c.Params("cropTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsCropType(database.DB, cropTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	crop := new(models.Crop)
	if err := c.Bind().Body(crop); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	crop.CropTypeId = cropTypeId
	crop, err = repository.CreateCrop(database.DB, crop)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(crop)
}

func DeleteCrop(c fiber.Ctx) error {
	cropTypeId, err := strconv.ParseUint(c.Params("cropTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsCropType(database.DB, cropTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err = repository.ExistsCrop(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	err = repository.DeleteCrop(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateCrop(c fiber.Ctx) error {
	cropTypeId, err := strconv.ParseUint(c.Params("cropTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsCropType(database.DB, cropTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err = repository.ExistsCropType(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	crop := new(models.Crop)
	if err := c.Bind().Body(crop); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	crop.Id = id
	crop, err = repository.UpdateCrop(database.DB, crop)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.JSON(crop)
}
