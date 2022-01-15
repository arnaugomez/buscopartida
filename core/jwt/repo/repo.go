package jwtRepo

import (
	"errors"
	"github.com/arnaugomez/buscopartida/core/env"
	"github.com/arnaugomez/buscopartida/core/jwt"
	jwtLib "github.com/dgrijalva/jwt-go"
	"time"
)

type repo struct {
	key []byte
}

func CreateRepo(envRepo env.Repo) jwt.Repo {
	return &repo{[]byte(envRepo.GetEnv().JwtSecretKey)}
}

func getExpirationTime() int64 {
	return time.Now().Add(time.Hour * 24 * 7).Unix()
}

func (r *repo) CreateJWT(userId uint) (string, error) {
	c := &claims{
		UserId: userId,
		StandardClaims: jwtLib.StandardClaims{
			ExpiresAt: getExpirationTime(),
		},
	}

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, c)

	return token.SignedString(r.key)
}

func (r *repo) GetUserIdByJWT(token string) (userId uint, err error) {
	c := &claims{}
	tkn, err := jwtLib.ParseWithClaims(token, c, func(t *jwtLib.Token) (interface{}, error) {
		return r.key, nil
	})
	if err != nil {
		return 0, err
	}
	if !tkn.Valid {
		err = errors.New("jwt token is invalid")
	}
	return c.UserId, err
}
