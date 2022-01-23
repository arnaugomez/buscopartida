package userRepo

import (
	"errors"
	"github.com/arnaugomez/buscopartida/core/common/utils/hashing"
	"github.com/arnaugomez/buscopartida/core/user"
	userDbModels "github.com/arnaugomez/buscopartida/core/user/repo/models"
	userDataTransformers "github.com/arnaugomez/buscopartida/core/user/repo/transformers"
	"gorm.io/gorm"
)

func (r repo) CreateUser(c *user.Credentials) (*user.User, error) {
	hash, err := hashing.Hash(c.Password)
	if err != nil {
		return nil, err
	}
	usr := userDataTransformers.CredentialsToUser(c, hash)
	r.db.Create(usr)
	return userDataTransformers.ToUserDomain(usr), nil
}

func (r repo) GetUserById(id uint) (*user.User, error) {
	u := &userDbModels.User{}
	err := r.db.Where(&userDbModels.User{ID: id}).Take(u).Error
	return userDataTransformers.ToUserDomain(u), err
}

func (r repo) GetUserByEmail(email string) (*user.User, error) {
	u := &userDbModels.User{}
	err := r.db.Where(&userDbModels.User{Email: email}).Take(u).Error
	return userDataTransformers.ToUserDomain(u), err
}

func (r repo) UserAlreadyExists(c *user.Credentials) bool {
	err := r.db.Where(&userDbModels.User{Name: c.Name}).Or(&userDbModels.User{Email: c.Email}).Take(&userDbModels.User{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (r repo) UpdatePassword(userId uint, password string) error {
	hash, err := hashing.Hash(password)
	if err != nil {
		return err
	}
	return r.db.Take(&userDbModels.User{}, userId).Update("password_hash", hash).Error
}

func (r repo) DeleteUserById(id uint) error {
	return r.db.Delete(&userDbModels.User{}, id).Error
}
