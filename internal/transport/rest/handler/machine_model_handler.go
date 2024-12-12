package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetMachineModels(c *fiber.Ctx) error {
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
	machineModels, err := repository.GetAllMachineModelsByMachineTypeId(database.DB, machineTypeId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	machineModelsResponse := make([]dto.MachineModelDto, 0, len(machineModels))
	for _, m := range machineModels {
		machineModelsResponse = append(machineModelsResponse, *mappers.FromMachineModelToDto(&m))
	}

	return c.JSON(machineModelsResponse)
}

func CreateMachineModel(c *fiber.Ctx) error {
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

	machineModelDto := new(dto.CreateMachineModelRequestDto)
	if err := c.BodyParser(machineModelDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	machine, err := repository.CreateMachineModel(database.DB, mappers.FromCreateRequestDtoToMachineModel(machineModelDto, machineTypeId))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromMachineModelToDto(machine))
}

func DeleteMachineModel(c *fiber.Ctx) error {
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

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err = repository.ExistsMachineModel(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	err = repository.DeleteMachineModel(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateMachineModel(c *fiber.Ctx) error {
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

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err = repository.ExistsMachineModel(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	machineModelDto := new(dto.UpdateMachineModelRequestDto)
	if err := c.BodyParser(machineModelDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	machineModel, err := repository.UpdateMachineModel(database.DB, id, mappers.FromUpdateRequestDtoToMachineModel(machineModelDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}

	return c.JSON(mappers.FromMachineModelToDto(machineModel))
}
