package di

import (
	"fmt"
	db2 "github.com/arnaugomez/buscopartida/deps/db"
	"github.com/arnaugomez/buscopartida/deps/env"
)

func Setup() Di {
	envRer := env.Setup()
	di := Di{
		envRer: envRer,
	}
	db := db2.Setup(di)
	fmt.Println(db)


	return di
}