package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetEquipmentTypes(c *fiber.Ctx) error {
	equipmentTypes, err := repository.GetAllEquipmentTypes(database.DB)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	equipmentTypesResponse := make([]dto.EquipmentTypeDto, 0, len(equipmentTypes))
	for _, m := range equipmentTypes {
		equipmentTypesResponse = append(equipmentTypesResponse, *mappers.FromEquipmentTypeToDto(&m))
	}

	return c.JSON(equipmentTypesResponse)
}

func CreateEquipmentType(c *fiber.Ctx) error {
	equipmentTypeDto := new(dto.CreateEquipmentTypeRequestDto)
	if err := c.BodyParser(equipmentTypeDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	equipmentType, err := repository.CreateEquipmentType(database.DB, mappers.FromCreateRequestDtoToEquipmentType(equipmentTypeDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromEquipmentTypeToDto(equipmentType))
}

func DeleteEquipmentType(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsEquipmentType(database.DB, id)
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

func UpdateEquipmentType(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsEquipmentType(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	equipmentTypeDto := new(dto.UpdateEquipmentTypeRequestDto)
	if err := c.BodyParser(equipmentTypeDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	equipmentType, err := repository.UpdateEquipmentType(database.DB, id, mappers.FromUpdateRequestDtoToEquipmentType(equipmentTypeDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.JSON(mappers.FromEquipmentTypeToDto(equipmentType))
}
