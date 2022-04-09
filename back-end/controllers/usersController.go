package controllers

import (
	"log"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/services"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var user models.User

	if err := fctx.BodyParser(&user); err != nil {
		log.Println("Error in method post CreateUser: ", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot trasnform JSON to struct: " + err.Error(),
		})
	}

	user.Password = services.SHA256Encoder(user.Password)

	err := db.Create(&user).Error
	if err != nil {
		log.Println("Error in method post CreateUser:", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot create user: " + err.Error(),
		})
	}

	return fctx.SendStatus(fiber.StatusCreated)

}
