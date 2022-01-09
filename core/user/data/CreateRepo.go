package userData

import (
	"github.com/arnaugomez/buscopartida/core/db"
	"github.com/arnaugomez/buscopartida/core/user"
	"github.com/arnaugomez/buscopartida/ctx"
)

func CreateRepo(database db.DB, context *ctx.Ctx) user.Repo {
	return repo{database, context}
}