package controllers

import (
	"v1/app/database"
	"v1/app/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// RegisterController handles user registration
func RegisterController(c *fiber.Ctx) error {
	// Parse request body
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	// Update user password with hashed password
	user.Password = string(hashedPassword)

	// Save the user to the database
	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user",
		})
	}

	// Clear sensitive information before sending the response
	user.Password = ""

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    user,
	})
}