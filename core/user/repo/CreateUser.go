package userRepo

import (
	"github.com/arnaugomez/buscopartida/core/common/utils/hashing"
	"github.com/arnaugomez/buscopartida/core/user"
	userDataTransformers "github.com/arnaugomez/buscopartida/core/user/repo/transformers"
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
