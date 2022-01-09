package ctxDomain

import (
	"github.com/arnaugomez/buscopartida/core/db"
	"github.com/arnaugomez/buscopartida/core/env/data"
	userData "github.com/arnaugomez/buscopartida/core/user/data"
	"github.com/arnaugomez/buscopartida/ctx"
)

func Setup() ctx.Ctx {
	var context ctx.Ctx
	envRepo := envData.CreateRepo()
	database := db.Setup(envRepo)
	userRepo := userData.CreateRepo(database, &context)
	context = ctx.Ctx{
		EnvRepo: envRepo,
		UserRepo: userRepo,
	}
	userRepo.GetUser("Hello");
	return context
}