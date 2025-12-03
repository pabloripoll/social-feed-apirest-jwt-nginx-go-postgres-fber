package database

import (
	userModel "apirest/domain/user/model"
	feedModel "apirest/domain/feed/model"

	"gorm.io/gorm"
)

// RunMigrations performs schema migrations using the provided gorm DB instance.
// Returns an error if AutoMigrate fails.
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&userModel.User{},
		&feedModel.Feed{},
	)
}


	// Migrate the schemas
	//err = DB.AutoMigrate(&model.Feed{}, &model.User{})
	//if err != nil {
	//	panic("Failed to migrate database schemas!")
	//}
	//fmt.Println("Database migrated.")