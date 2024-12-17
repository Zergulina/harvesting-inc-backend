package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository"
	"backend/internal/dto"
	"backend/internal/mappers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	posts, err := repository.GetAllPosts(database.DB)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	postsResponse := make([]dto.PostDto, 0, len(posts))
	for _, p := range posts {
		postsResponse = append(postsResponse, *mappers.FromPostToDto(&p))
	}

	return c.JSON(postsResponse)
}

func CreatePost(c *fiber.Ctx) error {
	postDto := new(dto.CreatePostRequestDto)
	if err := c.BodyParser(postDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	post, err := repository.CreatePost(database.DB, mappers.FromCreateRequestDtoToPost(postDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}

	return c.JSON(mappers.FromPostToDto(post))
}

func DeletePost(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsPost(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	err = repository.DeletePost(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.SendString("Успешно удалено")
}

func UpdatePost(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Неверный запрос")
	}
	isExist, err := repository.ExistsPost(database.DB, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка базы данных")
	}
	if !isExist {
		return c.Status(404).SendString("Не найдено")
	}
	postDto := new(dto.UpdatePostRequestDto)
	if err := c.BodyParser(postDto); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}
	post, err := repository.UpdatePost(database.DB, id, mappers.FromUpdateReqeustDtoToPost(postDto))
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления")
	}
	return c.JSON(mappers.FromPostToDto(post))
}
