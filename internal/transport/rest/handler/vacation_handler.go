package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetVacationsByPeopleId(c *fiber.Ctx) error {
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
	vacations, err := repository.GetAllVacationsByPeopleId(database.DB, peopleId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	vacationsResponse := make([]dto.VacationDto, 0, len(vacations))
	for _, v := range vacations {
		vacationsResponse = append(vacationsResponse, *mappers.FromVacationToDto(&v))
	}

	return c.JSON(vacationsResponse)
}

func CreateVacation(c *fiber.Ctx) error {
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

	vacationDto := new(dto.CreateVacationRequestDto)
	if err := c.BodyParser(vacationDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	vacation, err := repository.CreateVacation(database.DB, mappers.FromCreateRequestDtoToVacation(vacationDto, peopleId))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromVacationToDto(vacation))
}

func DeleteVacation(c *fiber.Ctx) error {
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

	vacationDto := new(dto.DeleteVacationRequestDto)
	if err := c.BodyParser(vacationDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	err = repository.DeleteVacation(database.DB, mappers.FromDeleteRequestDtoToVacation(vacationDto, peopleId))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}
