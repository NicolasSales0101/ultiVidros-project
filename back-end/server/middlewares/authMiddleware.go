package middlewares

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
)

func AuthRequired() func(fctx *fiber.Ctx) error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}
