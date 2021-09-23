package date_utils

import "time"

const (
	//2006-01-02 year month day
	//02-01-2006 day month year
	apiDateLayout = "02-01-2006T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
