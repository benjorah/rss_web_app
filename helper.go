package main

import (
	"regexp"
)

func removeHTMLTagsFromString(propertyString string) (sanitizedString string, err error) {

	regexObj, err := regexp.Compile(propertyString)

	if err != nil {

		return propertyString, err

	}

	return regexObj.ReplaceAllString("", "<[^>]*>"), err

}
