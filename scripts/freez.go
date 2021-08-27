package scripts

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

type Process struct {
	USER    string
	PID     string
	CPU     float32
	MEM     float32
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
	return data_slice
}
