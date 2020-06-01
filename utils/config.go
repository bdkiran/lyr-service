package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//loads in environment variables into the application
func init() {
	err := godotenv.Load()
	if err != nil {
		defaultLogger.Error.Fatal("Error loading .env file")
	}
}

//GetEnvVariableString reurns environment variables based on the key passed in
func GetEnvVariableString(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	errorString := fmt.Sprintf("Unable to find enviorment variable %s", key)
	return "", errors.New(errorString)
}
