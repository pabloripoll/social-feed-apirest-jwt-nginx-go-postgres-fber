package database

import (
	"fmt"
	"apirest/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB instance
var DB *gorm.DB

// Connect to the database
func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Europe/Madrid",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}
	fmt.Println("Connection opened to database")

	// Migrate the schemas
	err = DB.AutoMigrate(&model.FeedBook{}, &model.User{})
	if err != nil {
		panic("Failed to migrate database schemas!")
	}
	fmt.Println("Database migrated.")
}
