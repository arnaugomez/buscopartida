package di

import "github.com/arnaugomez/buscopartida/deps/env"

type Di struct {
	envRer env.Rer
}

func (di Di) EnvRer() env.Rer {
	return di.envRer
}
