package utils

import (
	"os"
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
func GetLogFileName() string {
	return "vegeta_" + Time() + ".log"
}
func CreateFile(path string) (string, error) {
	if ex := FileExists(path); ex {
		return path, nil
	}
	f, e := os.Create(path)
	if e != nil {
		return "", e
	}
	f.Close()
	return path, nil
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
