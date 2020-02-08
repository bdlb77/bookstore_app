package date_utils

import (
	"time"
)

const (
	apiDateString = "Y-m-dTH:i:sZ"
)

func GetNow() time.Time {
	return time.Now().UTC()
}
func GetNowString() string {
	return GetNow().Format(apiDateString)
}
