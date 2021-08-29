package scripts

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/kartik1998/vegeta/utils"
)

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

var ps_command = "ps aux | sort -rnk 4 | awk {'print $1\",\"$2\",\"$3\",\"$4\",\"$5\",\"$6\",\"$8\",\"$9\",\"$10\",\"$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$28,$30'} | sed '$d'"

func CollectProcessData(pid_set map[string]bool, m string, filepaths ...string) ([]string, error) {
	if m != "-1" {
		ps_command = ps_command + " | head -n " + m
	}
	out, err := exec.Command("bash", "-c", ps_command).Output()
	if err != nil {
		return []string{}, err
	}
	data_slice := cleanProcessData(string(out), pid_set)
	write_slice, _ := writeProcessData("./"+filepaths[0], data_slice, pid_set)
	return write_slice, nil
}

func writeProcessData(filepath string, values []string, pid_set map[string]bool) ([]string, error) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	var write_slice []string
	if err != nil {
		return []string{}, nil
	}
	defer file.Close()
	for i, value := range values {
		if value == "" {
			continue
		}
		value = utils.LogTime() + "," + value
		writestr := fmt.Sprintf("%s\n", value)
		if i == len(values)-1 && len(pid_set) == 0 {
			writestr = value
		}
		write_slice = append(write_slice, writestr)
		if _, err := file.WriteString(writestr); err != nil {
			log.Fatal(err)
		}
	}
	return write_slice, nil
}

func cleanProcessData(process_data string, pid_set map[string]bool) []string {
	data_slice := strings.Split(process_data, "\n")
	if len(pid_set) == 0 {
		return data_slice
	}
	var cleaned_slice []string
	for _, val := range data_slice {
		curr_slice := strings.Split(val, ",")
		if len(curr_slice) >= 2 && pid_set[curr_slice[1]] {
			cleaned_slice = append(cleaned_slice, val)
		}
	}
	return cleaned_slice
}
