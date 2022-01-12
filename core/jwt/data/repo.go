package jwtData

import (
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
	claims := &claims{
		UserId: userId,
		StandardClaims: jwtLib.StandardClaims{
			ExpiresAt: getExpirationTime(),
		},
	}

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)

	return token.SignedString(r.key)
}
