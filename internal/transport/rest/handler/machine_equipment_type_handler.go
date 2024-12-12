package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllEquipmentTypesByMachineTypeId(c *fiber.Ctx) error {
	machineTypeId, err := strconv.ParseUint(c.Params("machineTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsMachineType(database.DB, machineTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	equipmentTypes, err := repository.GetAllEquipmentTypesByMachineTypeId(database.DB, machineTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	equipmentTypesResponse := make([]dto.EquipmentTypeDto, 0, len(equipmentTypes))
	for _, e := range equipmentTypes {
		equipmentTypesResponse = append(equipmentTypesResponse, *mappers.FromEquipmentTypeToDto(&e))
	}

	return c.JSON(equipmentTypesResponse)
}

func CreateMachineEquipmentType(c *fiber.Ctx) error {
	machineTypeId, err := strconv.ParseUint(c.Params("machineTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsMachineType(database.DB, machineTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	equipmentTypeId, err := strconv.ParseUint(c.Params("equipmentTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err = repository.ExistsEquipmentType(database.DB, equipmentTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	isExist, err = repository.ExistsMachineEquipmentType(database.DB, machineTypeId, equipmentTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(400).SendString("Неверный запрос")
	}

	err = repository.CreateMachineEquipmentType(database.DB, machineTypeId, equipmentTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.SendString("Успешно создано")
}

func DeleteMachineEquipmentType(c *fiber.Ctx) error {
	machineTypeId, err := strconv.ParseUint(c.Params("machineTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsMachineType(database.DB, machineTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	equipmentTypeId, err := strconv.ParseUint(c.Params("equipmentTypeId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err = repository.ExistsEquipmentType(database.DB, equipmentTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	isExist, err = repository.ExistsMachineEquipmentType(database.DB, machineTypeId, equipmentTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(400).SendString("Неверный запрос")
	}

	err = repository.DeleteMachineEquipmentType(database.DB, machineTypeId, equipmentTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.SendString("Успешно удалено")
}
