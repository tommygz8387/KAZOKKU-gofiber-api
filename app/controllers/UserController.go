package controllers

import (
	"v1/app/database"
	"v1/app/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

// UserResponse defines the structure of the user response
type UserResponse struct {
	UserID     uint            `json:"user_id"`
	Name       string          `json:"name"`
	Email      string          `json:"email"`
	Address    string          `json:"address"`
	Photos     map[int]string  `json:"photos"`
	CreditCard CreditCardInfo  `json:"creditcard"`
}

// CreditCardInfo defines the structure of the credit card information in the response
type CreditCardInfo struct {
	Type    string `json:"type"`
	Number  string `json:"number"`
	Name    string `json:"name"`
	Expired string `json:"expired"`
	Cvv string `json:"cvv"`
}

func MaskCreditCardNumber(cardNumber string) string {
	if len(cardNumber) > 4 {
		masked := "***" + cardNumber[len(cardNumber)-4:]
		return masked
	}
	return cardNumber
}

func UserController(c *fiber.Ctx) error {
	var users []models.User

	// Fetch all users with their relations
	result := database.DB.Preload(clause.Associations).Omit("password", "CreatedAt", "UpdatedAt", "DeletedAt").Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	var userResponses []UserResponse
	for _, user := range users {
		userResponse := UserResponse{
			UserID:  user.ID,
			Email:   user.Email,   
			Address: user.Address, 
			Photos:  make(map[int]string),
			CreditCard: CreditCardInfo{
				Type:    user.CreditCard.Type,   
				Name:    user.CreditCard.Name,   
				Number:  MaskCreditCardNumber(user.CreditCard.Number),
				Expired: user.CreditCard.Expired,
				Cvv: user.CreditCard.Cvv,
			},
		}

		// Menyusun data foto
		for i, photo := range user.Photos {
			userResponse.Photos[i+1] = photo.Filename
		}

		userResponses = append(userResponses, userResponse)
	}

	return c.JSON(fiber.Map{
		"count": len(users),
		"rows": userResponses,
	})
}