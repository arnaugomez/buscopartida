package userDataTransformers

import (
	"github.com/arnaugomez/buscopartida/core/user"
	userDataModels "github.com/arnaugomez/buscopartida/core/user/repo/models"
)

func ToUserDomain(u *userDataModels.User) *user.User {
	return &user.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
}
