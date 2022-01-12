package userRepo

import (
	"github.com/arnaugomez/buscopartida/core/db"
	"github.com/arnaugomez/buscopartida/core/user"
	"github.com/arnaugomez/buscopartida/ctx"
)

// Implements user.Repo
type repo struct {
	db  db.DB
	ctx *ctx.Ctx
}

func CreateRepo(database db.DB, context *ctx.Ctx) user.Repo {
	return repo{database, context}
}
