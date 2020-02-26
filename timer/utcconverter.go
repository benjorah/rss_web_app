//Package timer contains methods manipulating date time
package timer

import (
	"fmt"
	"time"
)

//ConvertTimeToUTC converts a time string to UTC format
//It returns the time and any error encountered
func ConvertTimeToUTC(timeString string) (convertedTime *time.Time, err error) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05 -0700 MST", timeString)

	if err != nil {
		return nil, fmt.Errorf("[ERROR] convertTimeToUTC() in hepler : %s", err.Error())
	}

	convertedTime = &parsedTime
	return convertedTime, nil
}
