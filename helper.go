package main

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func removeHTMLTagsFromString(propertyString string) (sanitizedString string, err error) {

	regexObj, err := regexp.Compile(propertyString)

	if err != nil {

		return propertyString, err

	}

	return regexObj.ReplaceAllString("", "<[^>]*>"), err

}
