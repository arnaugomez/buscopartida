package db

import (
	"fmt"
	"github.com/arnaugomez/buscopartida/core/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Example struct {
	ID   uint
	Code string
}

// Gorm Database
type DB = *gorm.DB

func Setup(envRepo env.Repo) DB {
	env := envRepo.GetEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Madrid", env.Host, env.User, env.Password, env.DbName, env.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err2 := db.AutoMigrate(&Example{})
	if err2 != nil {
		panic(err)
	}

	return db
	// db.Create(&Example{Code: "D42"})
}
