package userDataModels

import "github.com/arnaugomez/buscopartida/core/db"

func RegisterModels(db db.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
