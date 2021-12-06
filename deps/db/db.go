package db

import (
	"fmt"
	"github.com/arnaugomez/buscopartida/deps/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Example struct {
	ID   uint
	Code string
}

type di interface {
	EnvRer() env.Rer
}

func Setup(di di) *gorm.DB {
	env := di.EnvRer().GetEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", env.Host, env.User, env.Password, env.DbName, env.Port)
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
