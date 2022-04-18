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

	// --------------------------------------------------------------------------------------------

	// General glasses routes
	glasses := main.Group("/glasses")
	glasses.Get("/single-glass/:id", controllers.ShowGlass)
	glasses.Get("/", controllers.ShowGlasses)
	glasses.Post("/", middlewares.AuthRequired(), middlewares.AdminAuthRequired(), controllers.CreateGlass)
	glasses.Put("/", middlewares.AuthRequired(), middlewares.AdminAuthRequired(), controllers.UpdateGlass)
	glasses.Delete("/delete-glass/:id", middlewares.AuthRequired(), middlewares.AdminAuthRequired(), controllers.DeleteGlass)

	// Common glasses routes
	commonGlss := glasses.Group("/common")
	commonGlss.Get("/", controllers.ShowCommonGlasses)

	// Tempered glasses routes
	temperedGlss := glasses.Group("/tempered")
	temperedGlss.Get("/", controllers.ShowTemperedGlasses)

	// --------------------------------------------------------------------------------------------

	// General users routes
	users := main.Group("/users")
	users.Post("/", controllers.CreateUser)

	// --------------------------------------------------------------------------------------------

	// General request and sales routes
	sales := main.Group("/sales")
	sales.Get("/single-sale/:id", controllers.ShowSale)
	sales.Get("/", controllers.ShowSales)
	sales.Post("/", controllers.CreateSale)
	sales.Put("/", controllers.UpdateSale)
	sales.Delete("/delete-sale/:id", controllers.DeleteSale)

	// --------------------------------------------------------------------------------------------

	// General parts routes
	parts := main.Group("/parts")
	parts.Get("/single-part/:id", controllers.ShowPart)
	parts.Get("/", controllers.ShowParts)
	parts.Post("/", controllers.CreatePart)
	parts.Put("/", controllers.UpdatePart)
	parts.Delete("/delete-part/:id", controllers.DeletePart)

	return router
}
