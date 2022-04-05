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

	// Common glasses routes
	commonGlss := glasses.Group("/common")
	commonGlss.Get("/", controllers.ShowCommonGlasses)

	// Tempered glasses routes
	temperedGlss := glasses.Group("/tempered")
	temperedGlss.Get("/", controllers.ShowTemperedGlasses)

	return router
}
