package userDataModels

import "time"

type User struct {
	ID           string
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string
	PasswordHash string
}
