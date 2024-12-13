package handler

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/helpers"
	"backend/internal/mappers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Register(c *fiber.Ctx) error {
	registerRequestDto := new(dto.RegisterRequestDto)
	if err := c.BodyParser(registerRequestDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	isExist, err := repository.ExistsPeopleByLogin(database.DB, registerRequestDto.Login)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}

	people := mappers.FromRegisterDtoToPeople(registerRequestDto)

	people, err = repository.CreatePeople(database.DB, people)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(people)
}

func Login(c *fiber.Ctx) error {
	login := new(dto.LoginRequestDto)
	if err := c.BodyParser(login); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	isExist, err := repository.ExistsPeopleByLogin(database.DB, login.Login)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(403).SendString("Неверный логин или пароль")
	}

	existingPeople, err := repository.GetPeopleByLogin(database.DB, login.Login)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	if existingPeople.PasswordHash != helpers.EncodeSha256(login.Login, config.DbSecretKey) {
		return c.Status(403).SendString("Неверный логин или пароль")
	}

	employees, err := repository.GetAllEmployeesByPeopleId(database.DB, existingPeople.Id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	postsClaim := ""
	for _, employee := range employees {
		post, err := repository.GetPostById(database.DB, employee.PostId)
		if err != nil {
			return c.Status(500).SendString("Ошибка базы данных")
		}
		postsClaim += post.Name + "|"
	}

	postsClaim = postsClaim[:len(postsClaim)-1]

	claims := jwt.MapClaims{
		"login": existingPeople.Login,
		"posts": postsClaim,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.JwtSecretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	login := claims["login"].(string)

	people, err := repository.GetPeopleByLogin(database.DB, login)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	employees, err := repository.GetAllEmployeesByPeopleId(database.DB, people.Id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	posts := make([]string, 0, 3)

	for _, employee := range employees {
		post, err := repository.GetPostById(database.DB, employee.PostId)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		posts = append(posts, post.Name)
	}

	return c.JSON(mappers.FromPeoplePostsToDto(people, posts))
}
