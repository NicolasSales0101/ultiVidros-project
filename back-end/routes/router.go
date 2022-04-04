package routes

import (
	"github.com/NicolasSales0101/ultiVidros-project/back-end/controllers"
	"github.com/gofiber/fiber/v2"
)

func ConfigRoutes(router *fiber.App) *fiber.App {
	main := router.Group("/api/v1")

	glasses := main.Group("/glasses")
	glasses.Get("/", controllers.ShowGlasses)

	return router
}
