package jwtData

import (
	"github.com/arnaugomez/buscopartida/core/env"
	"github.com/arnaugomez/buscopartida/core/jwt"
)

func CreateRepo(envRepo env.Repo) jwt.Repo {
	return repo{[]byte(envRepo.GetEnv().JwtSecretKey)}
}
