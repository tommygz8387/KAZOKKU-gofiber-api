package controllers

import (
	"v1/app/database"
	"v1/app/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// RegisterController handles user registration
func Register(c *fiber.Ctx) error {
	var requestData RegisterRequest
	if err := c.BodyParser(&requestData); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	newUser := models.User{
		Name: requestData.Name,
		Email:    requestData.Email,
		Password: string(hashedPassword),
		Address:  requestData.Address,
	}
	database.DB.Create(&newUser)

	photo := models.UserPhoto{
		UserID:   newUser.ID,
		Filename: requestData.Photo,
	}
	database.DB.Create(&photo)

	creditCard := models.UserCreditCard{
		UserID:  newUser.ID,
		Type:    requestData.CardType,
		Name:    requestData.CardName,
		Number:  requestData.CardNumber,
		Expired: requestData.CardExpired,
		Cvv:     requestData.CardCVV,
	}
	database.DB.Create(&creditCard)

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"user_id": newUser.ID,
	})
}