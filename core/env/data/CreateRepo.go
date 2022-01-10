package envData

import (
	"github.com/arnaugomez/buscopartida/core/env"
	"github.com/joho/godotenv"
	"os"
)

func CreateRepo() env.Repo {
	godotenv.Load(".env")
	envVars := env.Env{
		Host:         os.Getenv("HOST"),
		User:         os.Getenv("USER"),
		Password:     os.Getenv("PASSWORD"),
		DbName:       os.Getenv("DB_NAME"),
		Port:         os.Getenv("PORT"),
		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
	r := repo{envVars}

	return r
}
