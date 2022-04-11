package routes

import (
	"github.com/NicolasSales0101/ultiVidros-project/back-end/controllers"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/server/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ConfigRoutes(router *fiber.App) *fiber.App {

	// General main routes
	main := router.Group("/api/v1", logger.New())
	main.Post("/login", controllers.Login)

	// General glasses routes
	glasses := main.Group("/glasses")
	glasses.Get("/single-glass/:id", controllers.ShowGlass)
	glasses.Get("/", controllers.ShowGlasses)
	glasses.Post("/", controllers.CreateGlass)
	glasses.Put("/", controllers.UpdateGlass)
	glasses.Delete("/delete-glass/:id", middlewares.AuthRequired(), controllers.DeleteGlass)

	// Common glasses routes
	commonGlss := glasses.Group("/common")
	commonGlss.Get("/", controllers.ShowCommonGlasses)

	// Tempered glasses routes
	temperedGlss := glasses.Group("/tempered")
	temperedGlss.Get("/", controllers.ShowTemperedGlasses)

	// ---------------------------------------------------------

	// General users routes
	users := main.Group("/users")
	users.Post("/", controllers.CreateUser)

	return router
}
