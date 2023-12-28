package controllers

import (
	"v1/app/database"
	"v1/app/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Register Function handles user registration
func Register(c *fiber.Ctx) error {
    // Parse request body
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
        })
    }

    // Validate user data
    if err := user.Validate(); err != nil {
        var validationErrors []string
        if verr, ok := err.(validator.ValidationErrors); ok {
            for _, fieldError := range verr {
                validationErrors = append(validationErrors, fieldError.Error())
            }
        }
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Validation failed",
            "errors":  validationErrors,
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
            "message": result.Error.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "message": "User registered successfully",
        "user_id":    user.ID,
    })
}