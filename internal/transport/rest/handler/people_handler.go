package handler

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/helpers"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPeople(c *fiber.Ctx) error {
	posts := (helpers.ExtractPostsFromJwt(c))
	if !helpers.Contains(posts, config.AdminRole) && !helpers.Contains(posts, config.HrRole) {
		return c.Status(fiber.StatusUnauthorized).SendString("Недостаточно прав")
	}

	people, err := repository.GetAllPeople(database.DB)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	peopleResponse := make([]dto.PeopleDto, 0, len(people))
	for _, p := range people {
		peopleResponse = append(peopleResponse, *mappers.FromPeopleToDto(&p))
	}

	return c.JSON(peopleResponse)
}

func DeletePeople(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsPeople(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	err = repository.DeletePeople(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdatePeople(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsPeople(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	updatePeopleRequestDto := new(dto.UpdatePeopleRequestDto)
	if err := c.BodyParser(updatePeopleRequestDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	people, err := repository.UpdatePeople(database.DB, id, mappers.FromUpdatePeopleRequestDtoToPeople(updatePeopleRequestDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.JSON(mappers.FromPeopleToDto(people))
}
