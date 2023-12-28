package controllers

import (
	"v1/app/database"
	"v1/app/models"

	"github.com/gofiber/fiber/v2"
)

func UserController(c *fiber.Ctx) error {
	var users []models.User

	// Fetch all users with their relations
	result := database.DB.Preload("Photos").Preload("CreditCard").Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	// Clear sensitive information before sending the response
	for i := range users {
		users[i].Password = ""
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}