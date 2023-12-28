package models

import "gorm.io/gorm"

type UserCreditCard struct {
	gorm.Model
	UserID     uint
	User       User   // Belongs To relationship
	CardType   string   `gorm:"not null"`
	CardNumber string   `gorm:"not null"`
	CardName   string   `gorm:"not null"`
	CardCvv    string   `gorm:"not null"`
}
