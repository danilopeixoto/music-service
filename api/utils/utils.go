package utils

import (
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ReadSecret function
func ReadSecret(name string) string {
	filename := os.Getenv(name)
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(bytes))
}

// Validate function
func Validate(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}
