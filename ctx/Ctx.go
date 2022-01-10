package ctx

import (
	"github.com/arnaugomez/buscopartida/core/env"
	"github.com/arnaugomez/buscopartida/core/jwt"
	"github.com/arnaugomez/buscopartida/core/user"
)

type Ctx struct {
	EnvRepo  env.Repo
	UserRepo user.Repo
	JwtRepo  jwt.Repo
}
