package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetCrops(c *fiber.Ctx) error {
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
	crops, err := repository.GetAllCropsByCropTypeId(database.DB, cropTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	cropType, err := repository.GetCropTypeById(database.DB, cropTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	cropsResponse := make([]dto.CropDto, 0, len(crops))
	for _, crop := range crops {
		cropsResponse = append(cropsResponse, *mappers.FromCropToDto(&crop, cropType))
	}

	return c.JSON(cropsResponse)
}

func CreateCrop(c *fiber.Ctx) error {
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

	cropDto := new(dto.CreateCropRequestDto)
	if err := c.BodyParser(cropDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	crop, err := repository.CreateCrop(database.DB, mappers.FromCreateRequestDtoToCrop(cropDto, cropTypeId))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	cropType, err := repository.GetCropTypeById(database.DB, crop.CropTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromCropToDto(crop, cropType))
}

func DeleteCrop(c *fiber.Ctx) error {
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

func UpdateCrop(c *fiber.Ctx) error {
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

	cropDto := new(dto.UpdateCropRequestDto)
	if err := c.BodyParser(cropDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	crop, err := repository.UpdateCrop(database.DB, id, mappers.FromUpdateRequestDtoToCrop(cropDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}

	cropType, err := repository.GetCropTypeById(database.DB, crop.CropTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromCropToDto(crop, cropType))
}
