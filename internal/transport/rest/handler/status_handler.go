package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetStatuses(c *fiber.Ctx) error {
	statuses, err := repository.GetAllStatuses(database.DB)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	statusesResponse := make([]dto.StatusDto, 0, len(statuses))
	for _, s := range statuses {
		statusesResponse = append(statusesResponse, *mappers.FromStatusToDto(&s))
	}

	return c.JSON(statusesResponse)
}

func CreateStatus(c *fiber.Ctx) error {
	statusDto := new(dto.CreateStatusRequestDto)
	if err := c.BodyParser(statusDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	status, err := repository.CreateStatus(database.DB, mappers.FromCreateRequestDtoToStatus(statusDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromStatusToDto(status))
}

func DeleteStatus(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsStatus(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	err = repository.DeleteStatus(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateStatus(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsStatus(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	statusDto := new(dto.UpdateStatusRequestDto)
	if err := c.BodyParser(statusDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	status, err := repository.UpdateStatus(database.DB, id, mappers.FromUpdateRequestDtoToStatus(statusDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.JSON(mappers.FromStatusToDto(status))
}
