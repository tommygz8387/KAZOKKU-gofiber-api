package models

import "gorm.io/gorm"

type UserPhoto struct {
	gorm.Model
	UserID          uint
	User    		User // Belongs To relationship
	Filename string
}
