package app

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/transport/rest"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	config.InitEnv()

	if err := database.Connect(); err != nil {
		panic(err)
	}

	app := fiber.New()

	rest.RegisterRoutes(app)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.JwtSecretKey)},
	}))

	rest.RegisterProtectedRoutes(app)

	app.Listen(":3000")
}
