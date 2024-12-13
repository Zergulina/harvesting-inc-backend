package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetEquipments(c *fiber.Ctx) error {
	equipmentModelId, err := strconv.ParseUint(c.Params("equipmentModelId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsEquipmentModel(database.DB, equipmentModelId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	equipments, err := repository.GetAllEquipmentsByEquipmentModelId(database.DB, equipmentModelId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	equipmentsResponse := make([]dto.EquipmentDto, 0, len(equipments))
	for _, m := range equipments {
		equipmentsResponse = append(equipmentsResponse, *mappers.FromEquipmentToDto(&m))
	}

	return c.JSON(equipmentsResponse)
}

func CreateEquipment(c *fiber.Ctx) error {
	equipmentModelId, err := strconv.ParseUint(c.Params("equipmentModelId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsEquipmentModel(database.DB, equipmentModelId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	equipmentDto := new(dto.CreateEquipmentRequestDto)
	if err := c.BodyParser(equipmentDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	equipments, err := repository.CreateEquipment(database.DB, mappers.FromCreateRequestDtoToEquipment(equipmentDto, equipmentModelId))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromEquipmentToDto(equipments))
}

func DeleteEquipment(c *fiber.Ctx) error {
	equipmentModelId, err := strconv.ParseUint(c.Params("equipmentModelId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	invNumber, err := strconv.ParseUint(c.Params("invNumber"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}

	isExist, err := repository.ExistsEquipment(database.DB, equipmentModelId, invNumber)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	err = repository.DeleteEquipmentModel(database.DB, invNumber)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateEquipment(c *fiber.Ctx) error {
	equipmentModelId, err := strconv.ParseUint(c.Params("equipmentModelId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	invNumber, err := strconv.ParseUint(c.Params("invNumber"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}

	isExist, err := repository.ExistsEquipment(database.DB, equipmentModelId, invNumber)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	equipmentDto := new(dto.UpdateEquipmentRequestDto)
	if err := c.BodyParser(equipmentDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	equipmentModel, err := repository.UpdateEquipment(database.DB, equipmentModelId, invNumber, mappers.FromUpdateRequestDtoToEquipment(equipmentDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}

	return c.JSON(mappers.FromEquipmentToDto(equipmentModel))
}
