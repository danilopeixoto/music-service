package config

import (
	"os"

	"danilopeixoto.com/api/music/utils"
)

// APIConfig struct
type APIConfig struct {
	Port    string
	Version string
	LogPath string
}

// DatabaseConfig struct
type DatabaseConfig struct {
	Hostname     string
	Port         string
	DatabaseName string
	Username     string
	Password     string
}

// APIConfig instance
var apiConfig *APIConfig

// DatabaseConfig instance
var databaseConfig *DatabaseConfig

// Load function
func Load() (*APIConfig, *DatabaseConfig) {
	apiConfig = &APIConfig{
		Port:    os.Getenv("API_PORT"),
		Version: os.Getenv("API_VERSION"),
		LogPath: os.Getenv("API_LOG_PATH"),
	}

	databaseConfig = &DatabaseConfig{
		Hostname:     os.Getenv("DATABASE_HOSTNAME"),
		Port:         os.Getenv("DATABASE_PORT"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		Username:     utils.ReadSecret("DATABASE_USERNAME_FILE"),
		Password:     utils.ReadSecret("DATABASE_PASSWORD_FILE"),
	}

	return apiConfig, databaseConfig
}

// GetAPIConfig function
func GetAPIConfig() *APIConfig {
	return apiConfig
}

// GetDatabaseConfig function
func GetDatabaseConfig() *DatabaseConfig {
	return databaseConfig
}
