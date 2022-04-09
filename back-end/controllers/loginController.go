package controllers

import (
	"log"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/services"
	"github.com/gofiber/fiber/v2"
)

func Login(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var login models.Login
	err := fctx.BodyParser(&login)
	if err != nil {
		log.Println("Error in method Login:", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot bind JSON: " + err.Error(),
		})
	}

	var user models.User

	dbError := db.Where("email = ?", user.Email).First(&user).Error
	if dbError != nil {
		log.Println("Error in method Login:", dbError)
		return fctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "cannot find user: " + dbError.Error(),
		})
	}

	if user.Password != services.SHA256Encoder(user.Password) {
		log.Println("Error in verify credentials")
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		log.Println("Error in verify token:", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid token: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"token": token,
	})
}
