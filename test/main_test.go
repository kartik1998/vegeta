package test

import (
	"testing"

	"github.com/kartik1998/vegeta/utils"
)

func TestTime(t *testing.T) {
	T := utils.Time()
	if len(T) < 10 {
		t.Errorf("wrong length for time string" + T)
	}
}

func TestLogTime(t *testing.T) {
	T := utils.LogTime()
	if len(T) < 10 {
		t.Errorf("wrong length for time string" + T)
	}
}

func TestGetLogFileName(t *testing.T) {
	name := utils.GetLogFileName()[0:6]

	if name != "vegeta" {
		t.Errorf("log file name should start with vegeta" + name)
	}

}
