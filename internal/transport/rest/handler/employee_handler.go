package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetEmployees(c *fiber.Ctx) error {
	peopleId, err := strconv.ParseUint(c.Params("peopleId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsPeople(database.DB, peopleId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	employees, err := repository.GetAllEmployeesByPeopleId(database.DB, peopleId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	employeesResponse := make([]dto.EmployeeDto, 0, len(employees))
	for _, c := range employees {
		employeesResponse = append(employeesResponse, *mappers.FromEmployeeToDto(&c))
	}

	return c.JSON(employeesResponse)
}

func CreateEmployee(c *fiber.Ctx) error {
	peopleId, err := strconv.ParseUint(c.Params("peopleId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsPeople(database.DB, peopleId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	employeeDto := new(dto.CreateEmployeeRequestDto)
	if err := c.BodyParser(employeeDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	isExist, err = repository.ExistsPost(database.DB, employeeDto.PostId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	employee, err := repository.CreateEmployee(database.DB, mappers.FromCreateRequestDtoToEmployee(employeeDto, peopleId))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromEmployeeToDto(employee))
}

func DeleteEmployee(c *fiber.Ctx) error {
	peopleId, err := strconv.ParseUint(c.Params("peopleId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	postId, err := strconv.ParseUint(c.Params("postId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsEmployee(database.DB, peopleId, postId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	err = repository.DeleteEmployee(database.DB, peopleId, postId)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateEmployee(c *fiber.Ctx) error {
	peopleId, err := strconv.ParseUint(c.Params("peopleId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	postId, err := strconv.ParseUint(c.Params("postId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsEmployee(database.DB, peopleId, postId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	employeeDto := new(dto.UpdateEmployeeRequestDto)
	if err := c.BodyParser(employeeDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	employee, err := repository.UpdateEmployee(database.DB, peopleId, postId, mappers.FromUpdateRequestDtoToEmployee(employeeDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}

	return c.JSON(mappers.FromEmployeeToDto(employee))
}
