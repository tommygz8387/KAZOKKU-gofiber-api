package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name              string   `gorm:"not null" json:"name"`
	Email          	  string   `gorm:"unique;not null" validate:"required,min=8" json:"email"`
	Password          string   `gorm:"not null" validate:"required,min=8"`
	Address           string   `gorm:"not null" validate:"required" json:"address"`
	Photos            []UserPhoto
	CreditCard		  *UserCreditCard
}


func (u *User) Validate() error {
    validate := validator.New()
    return validate.Struct(u)
}