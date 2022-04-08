package controllers

import (
	"log"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"github.com/gofiber/fiber/v2"
)

func ShowCommonGlasses(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var commonGlasses []models.Glass

	//err := db.Where("category = ?", "common").Find(&commonGlasses).Error
	err := db.Find(&commonGlasses, "category = ?", "common").Error
	if err != nil {
		log.Println("Error in methot get ShowCommonGlasses:", err.Error())
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot find common glasses: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(commonGlasses)
}
