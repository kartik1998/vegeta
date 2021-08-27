package scripts

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// type Process struct {
// 	USER    string
// 	PID     int32
// 	CPU     float32
// 	MEM     float32
// 	VSZ     string
// 	RSS     string
// 	TT      string
// 	STARTED string
// 	TIME    string
// 	COMMAND string
// }

type Process struct {
	USER    string
	PID     string
	CPU     string
	MEM     string
	VSZ     string
	RSS     string
	TT      string
	STARTED string
	TIME    string
	COMMAND string
}

func CollectProcessData() []string {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return cleanProcessData(out.String())
}

func cleanProcessData(process_data string) []string {
	data_slice := strings.Split(process_data, "\n")
	// processed_data_slice := []Process{}
	for i := 1; i < len(data_slice); i++ {
		curr_slice := strings.Split(data_slice[i], " ")
		fmt.Println(curr_slice)
	}
	return data_slice
}
