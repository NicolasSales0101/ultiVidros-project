package controllers

import (
	"log"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"github.com/gofiber/fiber/v2"
)

func ShowTemperedGlasses(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var temperedGlasses []models.Glass

	err := db.Where("category = ?", "tempered").Find(&temperedGlasses).Error
	if err != nil {
		log.Println("Error in methot get ShowTemperedGlasses:", err.Error())
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot find tempered glasses: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(temperedGlasses)
}
