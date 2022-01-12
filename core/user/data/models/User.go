package userDataModels

import "time"

type User struct {
	ID        uint `gorm:"autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name         string
	Email        string
	PasswordHash string
}
