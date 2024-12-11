package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetCustomers(c *fiber.Ctx) error {
	customers, err := repository.GetAllCustomers(database.DB)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	customersResponse := make([]dto.CustomerDto, 0, len(customers))
	for _, c := range customers {
		customersResponse = append(customersResponse, *mappers.FromCustomerToDto(&c))
	}

	return c.JSON(customersResponse)
}

func CreateCustomer(c *fiber.Ctx) error {
	customerDto := new(dto.CreateCustomerRequestDto)
	if err := c.BodyParser(customerDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	customer, err := repository.CreateCustomer(database.DB, mappers.FromCreateRequestDtoToCustomer(customerDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromCustomerToDto(customer))
}

func DeleteCustomer(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}

	isExist, err := repository.ExistsCustomer(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	err = repository.DeleteCustomer(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateCustomer(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}

	isExist, err := repository.ExistsCustomer(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	customerDto := new(dto.UpdateCustomerRequestDto)
	if err := c.BodyParser(customerDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	employee, err := repository.UpdateCustomer(database.DB, id, mappers.FromUpdateRequestDtoToCustomer(customerDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}

	return c.JSON(mappers.FromCustomerToDto(employee))
}

func PatchCustomer(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}

	isExist, err := repository.ExistsCustomer(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	customerDto := new(dto.PatchCustomerRequestDto)
	if err := c.BodyParser(customerDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	employee, err := repository.UpdateCustomer(database.DB, id, mappers.FromPatchRequestDtoToCustomer(customerDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}

	return c.JSON(mappers.FromCustomerToDto(employee))
}
