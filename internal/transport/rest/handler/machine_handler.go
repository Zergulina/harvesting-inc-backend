package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetMachines(c *fiber.Ctx) error {
	machineModelId, err := strconv.ParseUint(c.Params("machineModelId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsMachineModel(database.DB, machineModelId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	machines, err := repository.GetAllMachinesByMachineModelId(database.DB, machineModelId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	machinesResponse := make([]dto.MachineDto, 0, len(machines))
	for _, m := range machines {
		machinesResponse = append(machinesResponse, *mappers.FromMachineToDto(&m))
	}

	return c.JSON(machinesResponse)
}

func CreateMachine(c *fiber.Ctx) error {
	machineModelId, err := strconv.ParseUint(c.Params("machineModelId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsMachineModel(database.DB, machineModelId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	machineDto := new(dto.CreateMachineRequestDto)
	if err := c.BodyParser(machineDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	machine, err := repository.CreateMachine(database.DB, mappers.FromCreateRequestDtoToMachine(machineDto, machineModelId))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromMachineToDto(machine))
}

func DeleteMachine(c *fiber.Ctx) error {
	machineModelId, err := strconv.ParseUint(c.Params("machineModelId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	invNumber, err := strconv.ParseUint(c.Params("invNumber"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}

	isExist, err := repository.ExistsMachine(database.DB, machineModelId, invNumber)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	err = repository.DeleteMachineModel(database.DB, invNumber)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateMachine(c *fiber.Ctx) error {
	machineModelId, err := strconv.ParseUint(c.Params("machineModelId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	invNumber, err := strconv.ParseUint(c.Params("invNumber"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}

	isExist, err := repository.ExistsMachine(database.DB, machineModelId, invNumber)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	machineDto := new(dto.UpdateMachineRequestDto)
	if err := c.BodyParser(machineDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	machineModel, err := repository.UpdateMachine(database.DB, machineModelId, invNumber, mappers.FromUpdateRequestDtoToMachine(machineDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}

	return c.JSON(mappers.FromMachineToDto(machineModel))
}
