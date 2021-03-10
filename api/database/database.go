package database

import (
	"fmt"

	"danilopeixoto.com/api/music/config"
	"danilopeixoto.com/api/music/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database instance
var instance *gorm.DB

// Connect function
func Connect() *gorm.DB {
	dbConfig := config.GetDatabaseConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbConfig.Hostname,
		dbConfig.Port,
		dbConfig.DatabaseName,
		dbConfig.Username,
		dbConfig.Password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database server.")
	}

	db.Exec("create extension if not exists \"uuid-ossp\"")
	db.AutoMigrate(models.GetDatabaseModels()...)

	instance = db

	return instance
}

// GetDatabase function
func GetDatabase() *gorm.DB {
	return instance
}
