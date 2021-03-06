package utils

import (
	"log"
	"os"
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func getCurrentTimeUTC() time.Time {
	return time.Now().UTC()
}

func Time() string {
	return getCurrentTimeUTC().Format(apiDateLayout)
}

func LogTime() string {
	return time.Now().Format("01-02-2006 15:04:05")
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

func WriteFileHeader(filepath string) {
	str := "LOG_TIME,USER,PID,CPU,MEM,VSZ,RSS,STAT,START,TIME,COMMAND"
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	if _, err := file.WriteString(str + "\n"); err != nil {
		log.Fatal(err)
	}
}
