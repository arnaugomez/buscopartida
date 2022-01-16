package userRepo

import (
	"github.com/arnaugomez/buscopartida/core/common/utils/hashing"
	userDbModels "github.com/arnaugomez/buscopartida/core/user/repo/models"
)

func (r repo) UpdatePassword(userId uint, password string) error {
	hash, err := hashing.Hash(password)
	if err != nil {
		return err
	}
	return r.db.Take(&userDbModels.User{}, userId).Update("password_hash", hash).Error
}
