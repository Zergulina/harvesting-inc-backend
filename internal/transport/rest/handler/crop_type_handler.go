package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetCropTypes(c *fiber.Ctx) error {
	cropTypes, err := repository.GetAllCropTypes(database.DB)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	cropTypesResponse := make([]dto.CropTypeDto, 0, len(cropTypes))
	for _, c := range cropTypes {
		cropTypesResponse = append(cropTypesResponse, *mappers.FromCropTypeToDto(&c))
	}

	return c.JSON(cropTypesResponse)
}

func CreateCropType(c *fiber.Ctx) error {
	cropTypeDto := new(dto.CreateCropTypeRequestDto)
	if err := c.BodyParser(cropTypeDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	cropType, err := repository.CreateCropType(database.DB, mappers.FromCreateRequestDtoToCropType(cropTypeDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromCropTypeToDto(cropType))
}

func DeleteCropType(c *fiber.Ctx) error {
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

func UpdateCropType(c *fiber.Ctx) error {
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
	cropTypeDto := new(dto.UpdateCropTypeRequestDto)
	if err := c.BodyParser(cropTypeDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	cropType, err := repository.UpdateCropType(database.DB, id, mappers.FromUpdateReqeustDtoToCropType(cropTypeDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.JSON(mappers.FromCropTypeToDto(cropType))
}
