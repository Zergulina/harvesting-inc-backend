package helpers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractPostsFromJwt(c *fiber.Ctx) []string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	posts := claims["posts"].(string)
	return strings.Split(posts, "|")
}
