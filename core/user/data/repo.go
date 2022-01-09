package userData

import (
	"fmt"
	"github.com/arnaugomez/buscopartida/core/db"
	"github.com/arnaugomez/buscopartida/core/user"
	"github.com/arnaugomez/buscopartida/ctx"
)

// Implements user.Repo
type repo struct {
	db db.DB
	ctx *ctx.Ctx
}

func (r repo) GetUser(name string) user.User  {
	fmt.Println(r.ctx)
	return user.User{}
}