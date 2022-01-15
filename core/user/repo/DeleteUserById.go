package userRepo

import (
	"github.com/arnaugomez/buscopartida/core/user/repo/models"
)

func (r repo) DeleteUserById(id uint) error {
	return r.db.Delete(&userDbModels.User{}, id).Error
}
