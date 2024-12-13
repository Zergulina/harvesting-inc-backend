package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetEquipmentModels(c *fiber.Ctx) error {
	equipmentTypeId, err := strconv.ParseUint(c.Params("equipmentTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsEquipmentType(database.DB, equipmentTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	equipmentModels, err := repository.GetAllEquipmentModelsByEquipmentTypeId(database.DB, equipmentTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	equipmentModelsResponse := make([]dto.EquipmentModelDto, 0, len(equipmentModels))
	for _, m := range equipmentModels {
		equipmentModelsResponse = append(equipmentModelsResponse, *mappers.FromEquipmentModelToDto(&m))
	}

	return c.JSON(equipmentModelsResponse)
}

func CreateEquipmentModel(c *fiber.Ctx) error {
	equipmentTypeId, err := strconv.ParseUint(c.Params("equipmentTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsEquipmentType(database.DB, equipmentTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	equipmentModelDto := new(dto.CreateEquipmentModelRequestDto)
	if err := c.BodyParser(equipmentModelDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	equipments, err := repository.CreateEquipmentModel(database.DB, mappers.FromCreateRequestDtoToEquipmentModel(equipmentModelDto, equipmentTypeId))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromEquipmentModelToDto(equipments))
}

func DeleteEquipmentModel(c *fiber.Ctx) error {
	equipmentTypeId, err := strconv.ParseUint(c.Params("equipmentTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsEquipmentType(database.DB, equipmentTypeId)
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
	isExist, err = repository.ExistsEquipmentModel(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	err = repository.DeleteEquipmentModel(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateEquipmentModel(c *fiber.Ctx) error {
	equipmentTypeId, err := strconv.ParseUint(c.Params("equipmentTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsEquipmentType(database.DB, equipmentTypeId)
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
	isExist, err = repository.ExistsEquipmentModel(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	equipmentModelDto := new(dto.UpdateEquipmentModelRequestDto)
	if err := c.BodyParser(equipmentModelDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	equipmentModel, err := repository.UpdateEquipmentModel(database.DB, id, mappers.FromUpdateRequestDtoToEquipmentModel(equipmentModelDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}

	return c.JSON(mappers.FromEquipmentModelToDto(equipmentModel))
}
