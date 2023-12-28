package controllers

import (
	"v1/app/database"
	"v1/app/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// RegisterRequest adalah struktur data untuk permintaan registrasi
        type RegisterRequest struct {
            Name       string `json:"name" validate:"required"`
            Email      string `json:"email" validate:"required,email"`
            Password   string `json:"password" validate:"required,min=6"`
            Address    string `json:"address" validate:"required"`
            Photo      string `json:"photo" validate:"required"`
            CardType   string `json:"card_type" validate:"required"`
            CardNumber string `json:"card_number" validate:"required"`
            CardName   string `json:"card_name" validate:"required"`
            CardExpired string `json:"card_expired" validate:"required"`
            CardCVV    string `json:"card_cvv" validate:"required"`
        }

// RegisterController handles user registration
func Register(c *fiber.Ctx) error {
	// Parse request body
	var requestData RegisterRequest
	if err := c.BodyParser(&requestData); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	// Create a new user
	newUser := models.User{
		Name: requestData.Name,
		Email:    requestData.Email,
		Password: string(hashedPassword),
		Address:  requestData.Address,
	}
	database.DB.Create(&newUser)

	// Attach the provided photo to the user
	photo := models.UserPhoto{
		UserID:   newUser.ID,
		Filename: requestData.Photo,
	}
	database.DB.Create(&photo)

	// Attach the provided credit card to the user
	creditCard := models.UserCreditCard{
		UserID:  newUser.ID,
		Type:    requestData.CardType,
		Name:    requestData.CardName,
		Number:  requestData.CardNumber,
		Expired: requestData.CardExpired,
		Cvv:     requestData.CardCVV,
	}
	database.DB.Create(&creditCard)

	// Clear sensitive information before sending the response
	newUser.Password = ""

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"user_id": newUser.ID,
	})
}