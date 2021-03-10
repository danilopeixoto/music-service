package log

import (
	"os"

	"danilopeixoto.com/api/music/config"
	"github.com/sirupsen/logrus"
)

// Log instance
var instance *logrus.Logger

// Initialize function
func Initialize() *logrus.Logger {
	apiConfig := config.GetAPIConfig()

	instance = logrus.New()

	file, err := os.OpenFile(
		apiConfig.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		instance.SetOutput(os.Stdout)
	} else {
		instance.SetOutput(file)
	}

	return instance
}

// GetLogger function
func GetLogger() *logrus.Logger {
	return instance
}
