package userData

import (
	cryptoUseCases "github.com/arnaugomez/buscopartida/core/crypto/useCases"
	"github.com/arnaugomez/buscopartida/core/user"
	userDataTransformers "github.com/arnaugomez/buscopartida/core/user/data/transformers"
)

func (r repo) CreateUser(c *user.Credentials) (*user.User, error) {
	hash, err := cryptoUseCases.Hash(c.Password)
	if err != nil {
		return nil, err
	}
	usr := userDataTransformers.CredentialsToUser(c, hash)
	r.db.Create(&usr)
	return userDataTransformers.ToUserDomain(usr), nil
}
