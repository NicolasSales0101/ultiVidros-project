package controllers

import "github.com/gofiber/fiber/v2"

func ShowGlasses(fctx *fiber.Ctx) error {
	return fctx.JSON(fiber.Map{
		"status": "ok",
	})

}
