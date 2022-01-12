package ctxDomain

import (
	"github.com/arnaugomez/buscopartida/core/db"
	"github.com/arnaugomez/buscopartida/core/env/data"
	jwtData "github.com/arnaugomez/buscopartida/core/jwt/data"
	userData "github.com/arnaugomez/buscopartida/core/user/repo"
	"github.com/arnaugomez/buscopartida/ctx"
)

func CreateCtx() *ctx.Ctx {
	var context = &ctx.Ctx{}
	envRepo := envData.CreateRepo()
	database := db.Setup(envRepo)
	userRepo := userData.CreateRepo(database, context)
	jwtRepo := jwtData.CreateRepo(envRepo)
	*context = ctx.Ctx{
		EnvRepo:  envRepo,
		UserRepo: userRepo,
		JwtRepo:  jwtRepo,
	}
	return context
}
