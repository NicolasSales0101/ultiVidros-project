package controllers

import (
	"log"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"github.com/gofiber/fiber/v2"
)

func ShowPart(fctx *fiber.Ctx) error {

	id := fctx.Params("id")

	db := database.GetDatabase()

	var part models.Part
	err := db.First(&part, "id = ?", id).Error
	if err != nil {
		log.Println("Error in method get ShowPart (specific id in url params): ", err.Error())
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot find part: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(part)
}

func ShowParts(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var parts []models.Part

	err := db.Find(&parts).Error
	if err != nil {
		log.Println("Error in method get ShowParts: ", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot list parts: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(parts)
}

func CreatePart(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var part models.Part

	if err := fctx.BodyParser(&part); err != nil {
		log.Println("Error in method post CreatePart: ", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot trasnform JSON to struct: " + err.Error(),
		})
	}

	err := db.Create(&part).Error
	if err != nil {
		log.Println("Error in method post CreatePart:", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot create part: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusCreated).JSON(part)

}

func UpdatePart(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var part models.Part

	if err := fctx.BodyParser(&part); err != nil {
		log.Println("Error in method post UpdatePart: ", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot trasnform JSON to struct: " + err.Error(),
		})
	}

	err := db.Omit("CreatedAt").Save(&part).Error
	if err != nil {
		log.Println("Error in method post CreatePart: ", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot update part: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusCreated).JSON(part)
}

func DeletePart(fctx *fiber.Ctx) error {

	id := fctx.Params("id")

	db := database.GetDatabase()

	err := db.Delete(&models.Part{}, "id = ?", id).Error
	if err != nil {
		log.Println("Error in method delete DeletePart:", err.Error())
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot delete part" + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Part removed succefully",
	})
}
