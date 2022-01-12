package userRepo

import (
	"github.com/arnaugomez/buscopartida/core/user"
	"github.com/arnaugomez/buscopartida/core/user/repo/models"
	userDataTransformers "github.com/arnaugomez/buscopartida/core/user/repo/transformers"
)

func (r repo) GetUserByEmail(email string) (*user.User, error) {
	u := &userDbModels.User{}
	err := r.db.Where(&userDbModels.User{Email: email}).Take(u).Error
	return userDataTransformers.ToUserDomain(u), err
}

