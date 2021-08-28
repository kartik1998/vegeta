package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kartik1998/vegeta/scripts"
	"github.com/kartik1998/vegeta/utils"
)

/*
* @params
* delayParam: -d, --delay
 */
var (
	delay         int64  = 5
	m             string = "-1"
	log_file_name string = utils.GetLogFileName()
)

func init() {
	for i := 1; i < len(os.Args)-1; i += 2 {
		param := os.Args[i]
		var err error
		if param == "-d" || param == "--delay" {
			delay, err = strconv.ParseInt(os.Args[i+1], 0, 8)
			check(err)
		} else if param == "-m" {
			_, e := strconv.ParseInt(os.Args[i+1], 0, 8)
			check(e)
			m = os.Args[i+1]
		} else {
			panic("invalid parameters passed")
		}
	}
}

func main() {
	createLogFiles()
	for t := range time.Tick(time.Duration(delay) * time.Second) {
		_, err := scripts.CollectProcessData(m, "./"+log_file_name)
		if err != nil {
			log.Fatal(fmt.Sprintf("%s, %s", t, err))
			os.Exit(1)
		}
	}
}

func createLogFiles() {
	filepath, e := utils.CreateFile("./" + log_file_name)
	check(e)
	utils.WriteFileHeader(filepath)
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
