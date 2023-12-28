package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name              string   `gorm:"not null"`
	Username          string   `gorm:"unique;not null"`
	Password          string   `gorm:"not null"`
	Address           string   `gorm:"not null"`
	Photos            []UserPhoto
	CreditCard		  *UserCreditCard
}
