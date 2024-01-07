package controllers

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