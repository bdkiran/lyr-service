package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//LoadEnvVariables loads in environment variables into the application
//This is needed only if you're working with a local .env file, for production, we are passing in our single env variable.
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		defaultLogger.Error.Println("Error loading .env file")
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
