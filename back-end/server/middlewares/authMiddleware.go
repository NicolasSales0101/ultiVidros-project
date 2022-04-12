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

func AdminAuthRequired() fiber.Handler {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	adminLevel := os.Getenv("ADMIN_LEVEL")
	adminPass := os.Getenv("ADMIN_PASSWORD")

	return func(fctx *fiber.Ctx) error {
		headers := fctx.GetReqHeaders()
		for i, v := range headers {
			if i == "AdminAuthLevel" {
				if v != adminLevel {
					return fctx.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
						"error": "Unauthorized",
					})
				}
			}

			if i == "AdminAuthPass" {
				if v != adminPass {
					return fctx.Status(fiber.StatusNetworkAuthenticationRequired).JSON(fiber.Map{
						"error": "Unathorized",
					})
				}
			}
		}

		return fctx.Next()
	}
}
