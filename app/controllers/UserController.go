package controllers

import (
	"v1/app/database"
	"v1/app/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

type UpdateUserRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email" validate:"email"`
	Password   string `json:"password" validate:"min=6"`
	Address    string `json:"address"`
	Photo      string `json:"photo"`
	CardType   string `json:"card_type"`
	CardNumber string `json:"card_number"`
	CardName   string `json:"card_name"`
	CardExpired string `json:"card_expired"`
	CardCVV    string `json:"card_cvv"`
}

func MaskCreditCardNumber(cardNumber string) string {
	if len(cardNumber) > 4 {
		masked := "***" + cardNumber[len(cardNumber)-4:]
		return masked
	}
	return cardNumber
}

func GetUserList(c *fiber.Ctx) error {
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

func GetUserById(c *fiber.Ctx) error {
	var user models.User
	id := c.Params("id")
	result := database.DB.Preload(clause.Associations).Omit("password", "CreatedAt", "UpdatedAt", "DeletedAt").First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	
	// Prepare the user response
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
			Cvv:     user.CreditCard.Cvv,
		},
	}

	for i, photo := range user.Photos {
		userResponse.Photos[i+1] = photo.Filename
	}

	// Return the user response
	return c.JSON(fiber.Map{
		"user": userResponse,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	id := c.Params("id")
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	
	// Parse the request body into requestData
	var requestData UpdateUserRequest
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing request data: " + err.Error(),
		})
	}

	// Update user fields if provided in the request
	if requestData.Name != "" {
		user.Name = requestData.Name
	}
	if requestData.Email != "" {
		user.Email = requestData.Email
	}
	if requestData.Address != "" {
		user.Address = requestData.Address
	}
	if len(requestData.Password) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to hash password: " + err.Error(),
			})
		}
		user.Password = string(hashedPassword)
	}

	// Save the updated user to the database
	saveResult := database.DB.Save(&user)
	if saveResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user: " + saveResult.Error.Error(),
		})
	}

	// Return the updated user information
	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"user_id": user.ID,
	})
}
