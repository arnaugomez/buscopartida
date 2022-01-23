package userDbModels

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID uint
	User *User
	Slug string
	FirstName string
	LastName string
	Description string
	ProfileImg string
	BackgroundImg string
}