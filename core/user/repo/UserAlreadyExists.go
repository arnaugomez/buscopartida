package userRepo

import (
	"errors"
	"github.com/arnaugomez/buscopartida/core/user"
	userDbModels "github.com/arnaugomez/buscopartida/core/user/repo/models"
	"gorm.io/gorm"
)

func (r repo) UserAlreadyExists(c *user.Credentials) bool {
	err := r.db.Where(&userDbModels.User{Name: c.Name}).Or(&userDbModels.User{Email: c.Email}).Take(&userDbModels.User{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}