package userDataModels

type autoMigrator interface {
	AutoMigrate(model ...interface{}) error
}

func RegisterModels(db autoMigrator) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
