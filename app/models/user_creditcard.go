package models

import "gorm.io/gorm"

type UserCreditCard struct {
	gorm.Model `json:"-"`
	UserID     uint		`gorm:"unique;not null" json:"-"`
	User       User   	`json:"-"` // Belongs To relationship
	Type   string   `gorm:"not null"`
	Number string   `gorm:"not null"`
	Name   string   `gorm:"not null"`
	Expired   string   `gorm:"not null"`
	Cvv    string   `gorm:"not null"`
}
