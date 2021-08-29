package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kartik1998/vegeta/graph"
	"github.com/kartik1998/vegeta/scripts"
	"github.com/kartik1998/vegeta/utils"
)

/*
* @params
* delayParam: -d, --delay
* top M processes log counter: -m
* processId to trace: -p
* to generate a graph use: --graph
 */
var (
	delay          int64           = 5
	m              string          = "-1"
	log_file_name  string          = utils.GetLogFileName()
	pid_set        map[string]bool = make(map[string]bool)
	generate_graph bool            = false
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
		} else if param == "-p" {
			pid_set[os.Args[i+1]] = true
		} else {
			panic("invalid parameters passed")
		}
	}
	for i := 1; i < len(os.Args); i += 1 {
		if os.Args[i] == "--graph" || os.Args[i] == "-g" {
			generate_graph = true
		}
	}
}

func main() {
	var x_axis []string
	var y_axis []string
	createLogFiles()
	for t := range time.Tick(time.Duration(delay) * time.Second) {
		write_slice, err := scripts.CollectProcessData(pid_set, m, "./"+log_file_name)
		if generate_graph {
			slice := strings.Split(write_slice[0], ",")
			x_axis = append(x_axis, slice[0])
			y_axis = append(y_axis, slice[4])
			graph.GenerateGraph(log_file_name, x_axis, y_axis)
		}
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
