package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetMachineTypes(c *fiber.Ctx) error {
	machineTypes, err := repository.GetAllMachineTypes(database.DB)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	machineTypesResponse := make([]dto.MachineTypeDto, 0, len(machineTypes))
	for _, m := range machineTypes {
		machineTypesResponse = append(machineTypesResponse, *mappers.FromMachineTypeToDto(&m))
	}

	return c.JSON(machineTypesResponse)
}

func CreateMachineType(c *fiber.Ctx) error {
	machineTypeDto := new(dto.CreateMachineTypeRequestDto)
	if err := c.BodyParser(machineTypeDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	machineType, err := repository.CreateMachineType(database.DB, mappers.FromCreateRequestDtoToMachineType(machineTypeDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromMachineTypeToDto(machineType))
}

func DeleteMachineType(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsMachineType(database.DB, id)
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

func UpdateMachineType(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsMachineType(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	machineTypeDto := new(dto.UpdateMachineTypeRequestDto)
	if err := c.BodyParser(machineTypeDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	machineType, err := repository.UpdateMachineType(database.DB, id, mappers.FromUpdateRequestDtoToMachineType(machineTypeDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.JSON(mappers.FromMachineTypeToDto(machineType))
}
