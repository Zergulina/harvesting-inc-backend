package app

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/transport/rest"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v3"
)

func Run() {
	config.InitEnv()

	if err := database.Connect(); err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.JwtSecretKey)},
	}))

	rest.RegisterRoutes(app)

	app.Listen(":3000")
}
