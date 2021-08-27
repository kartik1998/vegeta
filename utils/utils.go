package utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func getCurrentTime() time.Time {
	return time.Now().UTC()
}
func Time() string {
	return getCurrentTime().Format(apiDateLayout)
}
