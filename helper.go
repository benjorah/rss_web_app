package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

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

func convertTimeToUTC(timeString string) (convertedTime *time.Time, err error) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05 -0700 MST", timeString)

	if err != nil {
		return nil, fmt.Errorf("[ERROR] convertTimeToUTC() in hepler : %s", err.Error())
	}

	convertedTime = &parsedTime
	return convertedTime, nil
}
