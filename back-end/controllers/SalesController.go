package controllers

import (
	"log"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"github.com/gofiber/fiber/v2"
)

func ShowSales(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var sales []models.Sale

	err := db.Find(&sales).Error
	if err != nil {
		log.Println("Error in method Get ShowSales:", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot list sales: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(sales)
}

func ShowSale(fctx *fiber.Ctx) error {
	id := fctx.Params("id")

	db := database.GetDatabase()

	var sale models.Sale

	err := db.First(&sale, "id = ?", id).Error
	if err != nil {
		log.Println("Error in method Get ShowSale (specific id in url params):", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot find sale: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(sale)
}

func CreateSale(fctx *fiber.Ctx) error {
	var sale models.Sale

	if err := fctx.BodyParser(&sale); err != nil {
		log.Println("Error in method Post CreateSale:", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot create sale: " + err.Error(),
		})
	}

	db := database.GetDatabase()
	err := db.Create(&sale).Error
	if err != nil {
		log.Println("Error in method Post CreateSale:", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot create sale: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusCreated).JSON(sale)
}

func UpdateSale(fctx *fiber.Ctx) error {
	var sale models.Sale

	if err := fctx.BodyParser(&sale); err != nil {
		log.Println("Error in method Put UpdateSale:", err)
		return fctx.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error": "cannot update sale: " + err.Error(),
		})
	}

	db := database.GetDatabase()

	err := db.Omit("CreatedAt").Save(&sale).Error
	if err != nil {
		log.Println("Error in method Put UpdateSale:", err)
		return fctx.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error": "cannot update sale: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(sale)
}

func DeleteSale(fctx *fiber.Ctx) error {
	id := fctx.Params("id")

	db := database.GetDatabase()

	err := db.Delete(&models.Sale{}, "id = ?", id).Error
	if err != nil {
		log.Println("Error in method Delete DeleteSale:", err)
		return fctx.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error": "cannot delete sale: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Sale removed successfully",
	})
}
