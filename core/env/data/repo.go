package envData

import "github.com/arnaugomez/buscopartida/core/env"

type repo struct {
	env env.Env
}

func (r repo) GetEnv() env.Env {
	return r.env
}