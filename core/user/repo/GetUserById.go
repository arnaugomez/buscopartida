package userRepo

import (
	"github.com/arnaugomez/buscopartida/core/user"
	"github.com/arnaugomez/buscopartida/core/user/repo/models"
	userDataTransformers "github.com/arnaugomez/buscopartida/core/user/repo/transformers"
)

func (r repo) GetUserById(id uint) (*user.User, error) {
	u := &userDbModels.User{}
	err := r.db.Where(&userDbModels.User{ID: id}).Take(u).Error
	return userDataTransformers.ToUserDomain(u), err
}

