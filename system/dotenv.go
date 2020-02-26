//Package system contains functions that makes it easier to work with the system OS by abstracting functionality
package system

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//GetEnvVariable returns an environment variable
func GetEnvVariable(key string) (variable string) {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
