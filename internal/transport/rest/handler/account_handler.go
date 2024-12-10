package handler

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/helpers"
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func Register(c fiber.Ctx) error {
	people := new(models.People)
	if err := c.Bind().Body(people); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	isExist, err := repository.ExistsPeopleByLogin(database.DB, people.Login)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	people.PasswordHash = helpers.EncodeSha256(people.PasswordHash, config.DbSecretKey)

	people, err = repository.CreatePeople(database.DB, people)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(people)
}

func Login(c fiber.Ctx) error {
	people := new(models.People)
	if err := c.Bind().Body(people); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	isExist, err := repository.ExistsPeopleByLogin(database.DB, people.Login)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(403).SendString("Неверный логин или пароль")
	}

	existingPeople, err := repository.GetPeopleByLogin(database.DB, people.Login)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	if existingPeople.PasswordHash != helpers.EncodeSha256(people.Login, config.DbSecretKey) {
		return c.Status(403).SendString("Неверный логин или пароль")
	}

	employees, err := repository.GetAllEmployeesByPeopleId(database.DB, people.Id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	postsClaim := "|"
	for _, employee := range employees {
		post, err := repository.GetPostById(database.DB, employee.PostId)
		if err != nil {
			return c.Status(500).SendString("Ошибка базы данных")
		}
		postsClaim += post.Name + "|"
	}

	claims := jwt.MapClaims{
		"login": existingPeople.Login,
		"posts": postsClaim,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.JwtSecretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
