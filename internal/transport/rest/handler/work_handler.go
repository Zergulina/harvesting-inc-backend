package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetWorksByCustomerId(c *fiber.Ctx) error {
	customerId, err := strconv.ParseUint(c.Params("customerId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsCustomer(database.DB, customerId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	works, err := repository.GetAllWorksByCustomerId(database.DB, customerId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	worksResponse := make([]dto.WorkDto, 0, len(works))
	for _, w := range works {
		worksResponse = append(worksResponse, *mappers.FromWorkToDto(&w))
	}

	return c.JSON(worksResponse)
}

func CreateWork(c *fiber.Ctx) error {
	fieldId, err := strconv.ParseUint(c.Params("fieldId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsField(database.DB, fieldId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	workDto := new(dto.CreateWorkRequestDto)
	if err := c.BodyParser(workDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	work, err := repository.CreateWork(database.DB, mappers.FromCreateRequestDtoToWork(workDto, fieldId))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromWorkToDto(work))
}

func DeleteWork(c *fiber.Ctx) error {
	fieldId, err := strconv.ParseUint(c.Params("fieldId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsField(database.DB, fieldId)
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
	isExist, err = repository.ExistsWork(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	err = repository.DeleteWork(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateWork(c *fiber.Ctx) error {
	fieldId, err := strconv.ParseUint(c.Params("customerId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsField(database.DB, fieldId)
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
	isExist, err = repository.ExistsWork(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	workDto := new(dto.UpdateWorkRequestDto)
	if err := c.BodyParser(workDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	work, err := repository.UpdateWork(database.DB, id, mappers.FromUpdateRequestDtoToWork(workDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}

	return c.JSON(mappers.FromWorkToDto(work))
}
