package database

import (
	"fmt"
	"os"

	"danilopeixoto.com/api/music/models"
	"danilopeixoto.com/api/music/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database instance
var instance *gorm.DB

// Connect function
func Connect() {
	hostname := os.Getenv("DATABASE_HOSTNAME")
	port := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")

	username := utils.ReadSecret("DATABASE_USERNAME_FILE")
	password := utils.ReadSecret("DATABASE_PASSWORD_FILE")

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		hostname, port, databaseName, username, password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database server.")
	}

	db.Exec("create extension if not exists \"uuid-ossp\"")
	db.AutoMigrate(models.GetDatabaseModels()...)

	instance = db
}

// GetDatabase function
func GetDatabase() *gorm.DB {
	return instance
}
