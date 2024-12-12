package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetWorkTripsByWorkId(c *fiber.Ctx) error {
	workId, err := strconv.ParseUint(c.Params("workId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsWork(database.DB, workId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	workTrips, err := repository.GetAllWorkTripsByWorkId(database.DB, workId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	workTripsResponse := make([]dto.WorkTripDto, 0, len(workTrips))
	for _, w := range workTrips {
		workTripsResponse = append(workTripsResponse, *mappers.FromWorkTripToDto(&w))
	}

	return c.JSON(workTripsResponse)
}

func CreateWorkTrip(c *fiber.Ctx) error {
	workId, err := strconv.ParseUint(c.Params("workId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsField(database.DB, workId)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	workTripDto := new(dto.CreateWorkTripRequestDto)
	if err := c.BodyParser(workTripDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	workTrip, err := repository.CreateWorkTrip(database.DB, mappers.FromCreateRequestDtoToWorkTrip(workTripDto, workId))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromWorkTripToDto(workTrip))
}

func DeleteWorkTrip(c *fiber.Ctx) error {
	workId, err := strconv.ParseUint(c.Params("workId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsWork(database.DB, workId)
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
	isExist, err = repository.ExistsWorkTrip(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	err = repository.DeleteWorkTrip(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdateWorkTrip(c *fiber.Ctx) error {
	workId, err := strconv.ParseUint(c.Params("workId"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsWork(database.DB, workId)
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
	isExist, err = repository.ExistsWorkTrip(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	workTripDto := new(dto.UpdateWorkTripRequestDto)
	if err := c.BodyParser(workTripDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	workTrip, err := repository.UpdateWorkTrip(database.DB, id, mappers.FromUpdateRequestDtoToWorkTrip(workTripDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}

	return c.JSON(mappers.FromWorkTripToDto(workTrip))
}
