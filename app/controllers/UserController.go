package controllers

import (
	"fmt"
	"strconv"
	"v1/app/database"
	"v1/app/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

func GetUserList(c *fiber.Ctx) error {
	var users []models.User
	const defaultSortBy = "id"
	const defaultSortOrder = "ASC"

	sortByQuery := c.Query("sb", defaultSortBy)
	sortOrderQuery := c.Query("ob", defaultSortOrder)

	if sortOrderQuery != "ASC" && sortOrderQuery != "DESC" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid sort order value",
		})
	}

	orderQuery := fmt.Sprintf("%s %s", sortByQuery, sortOrderQuery)

	const defaultLimit = 10

	limitQuery := c.Query("lt", fmt.Sprint(defaultLimit))

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid limit value",
		})
	}

	result := database.DB.Preload(clause.Associations).
		Omit("password", "CreatedAt", "UpdatedAt", "DeletedAt").
		Order(orderQuery).
		Limit(limit).
		Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching users",
		})
	}
	
	var userResponses []UserResponse
	for _, user := range users {
		userResponse := UserResponse{
			UserID:  user.ID,
			Name:   user.Name,   
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
	
	var requestData UpdateUserRequest
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing request data: " + err.Error(),
		})
	}

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

	saveResult := database.DB.Save(&user)
	if saveResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user: " + saveResult.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"user_id": user.ID,
	})
}