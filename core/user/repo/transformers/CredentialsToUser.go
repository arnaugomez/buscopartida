package userDataTransformers

import (
	"github.com/arnaugomez/buscopartida/core/user"
	userDataModels "github.com/arnaugomez/buscopartida/core/user/repo/models"
)

func CredentialsToUser(c *user.Credentials, passwordHash string) *userDataModels.User {
	return &userDataModels.User{
		Name:         c.Name,
		Email:        c.Email,
		PasswordHash: passwordHash,
	}
}
