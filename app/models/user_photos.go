package models

import "gorm.io/gorm"

type UserPhoto struct {
	gorm.Model `json:"-"`
	UserID          uint `json:"-"`
	User    		User `json:"-"` // Belongs To relationship
	Filename string
}
