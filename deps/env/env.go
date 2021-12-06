package env

import (
	"github.com/joho/godotenv"
	"os"
)

func Setup() Rer {
	godotenv.Load(".env")
	env := Env{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("PORT"),
	}
	r := envRepository{env}

	return r
}
