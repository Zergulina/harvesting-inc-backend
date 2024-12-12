package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetFields(c *fiber.Ctx) error {
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

	crops, err := repository.GetAllFieldsByCustomerId(database.DB, customerId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	cropsResponse := make([]dto.FieldDto, 0, len(crops))
	for _, c := range crops {
		cropsResponse = append(cropsResponse, *mappers.FromFieldToDto(&c))
	}

	return c.JSON(cropsResponse)
}

func CreateField(c *fiber.Ctx) error {
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

	fieldDto := new(dto.CreateFieldRequestDto)
	if err := c.BodyParser(fieldDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	field, err := repository.CreateField(database.DB, mappers.FromCreateRequestToField(fieldDto, customerId))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromFieldToDto(field))
}

func DeleteField(c *fiber.Ctx) error {
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

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err = repository.ExistsField(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	err = repository.DeleteField(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateField(c *fiber.Ctx) error {
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

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err = repository.ExistsField(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	fieldDto := new(dto.UpdateFieldRequestDto)
	if err := c.BodyParser(fieldDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	field, err := repository.UpdateField(database.DB, id, mappers.FromUpdateRequestToField(fieldDto, customerId))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}

	return c.JSON(mappers.FromFieldToDto(field))
}
