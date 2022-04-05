package routes

import (
	"github.com/NicolasSales0101/ultiVidros-project/back-end/controllers"
	"github.com/gofiber/fiber/v2"
)

func ConfigRoutes(router *fiber.App) *fiber.App {
	main := router.Group("/api/v1")

	// General glasses routes
	glasses := main.Group("/glasses")
	glasses.Get("/:id", controllers.ShowGlass)
	glasses.Get("/", controllers.ShowGlasses)
	glasses.Post("/", controllers.CreateGlass)
	glasses.Put("/", controllers.UpdateGlass)
	glasses.Delete("/:id", controllers.DeleteGlass)

	return router
}
