package middlewares

import (
	"github.com/NicolasSales0101/ultiVidros-project/back-end/services"
	"github.com/gofiber/fiber/v2"
)

// Attention here - In Progress

func Auth() fiber.Handler {

	return func(fctx *fiber.Ctx) error {

		const BearerSchema = "Bearer "
		header := fctx.GetReqHeaders()
		for _, v := range header {
			if v == "" {
				return fctx.SendStatus(fiber.StatusNetworkAuthenticationRequired)
			}
		}

		var token string

		for _, v := range header {
			token = v[len(BearerSchema):]
		}

		if !services.NewJWTService().ValidateToken(token) {
			return fctx.SendStatus(fiber.StatusNetworkAuthenticationRequired)
		}

		return nil

	}
}
