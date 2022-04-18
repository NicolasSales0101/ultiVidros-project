package controllers

import (
	"log"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"github.com/gofiber/fiber/v2"
)

func ShowGlass(fctx *fiber.Ctx) error {

	id := fctx.Params("id")

	db := database.GetDatabase()

	var glass models.Glass
	err := db.First(&glass, "id = ?", id).Error
	if err != nil {
		log.Println("Error in method get ShowGlass (specific id in url params): ", err.Error())
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot find glass: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(glass)
}

func ShowGlasses(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var glasses []models.Glass

	err := db.Find(&glasses).Error
	if err != nil {
		log.Println("Error in method get ShowGlasses: ", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot list glasses: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(glasses)
}

func CreateGlass(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var glass models.Glass

	if err := fctx.BodyParser(&glass); err != nil {
		log.Println("Error in method post CreateGlass: ", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot trasnform JSON to struct: " + err.Error(),
		})
	}

	err := db.Create(&glass).Error
	if err != nil {
		log.Println("Error in method post CreateGlass:", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot create glass: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusCreated).JSON(glass)

}

func UpdateGlass(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var glass models.Glass

	if err := fctx.BodyParser(&glass); err != nil {
		log.Println("Error in method post UpdateGlass: ", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot trasnform JSON to struct: " + err.Error(),
		})
	}

	err := db.Omit("CreatedAt").Save(&glass).Error
	if err != nil {
		log.Println("Error in method post CreateGlass: ", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot update glass: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusCreated).JSON(glass)
}

func DeleteGlass(fctx *fiber.Ctx) error {

	id := fctx.Params("id")

	db := database.GetDatabase()

	err := db.Delete(&models.Glass{}, "id = ?", id).Error
	if err != nil {
		log.Println("Error in method delete DeleteGlass:", err.Error())
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot delete glass" + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Glass removed succefully",
	})
}
