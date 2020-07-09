package date_utils

import "time"

const (
	apiDateLayout = "2020-06-29T17:24:30Z"
	apiDbLayout = "2020-06-29 17:24:30"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat()string  {
	return GetNow().Format(apiDbLayout)
}