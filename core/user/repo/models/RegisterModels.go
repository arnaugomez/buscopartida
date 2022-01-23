package userDbModels

type autoMigrator interface {
	AutoMigrate(model ...interface{}) error
}

func RegisterModels(db autoMigrator) error {
	err := db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		return err
	}
	return nil
}
