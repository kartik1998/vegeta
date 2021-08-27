package scripts

import (
	"os/exec"
	"strings"
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

const ps_command = "ps aux | head -1 | awk {'print $1\",\"$2\",\"$3\",\"$4\",\"$5\",\"$6\",\"$8\",\"$9\",\"$10\",\"$11'}; ps aux | sort -rnk 4 | awk {'print $1\",\"$2\",\"$3\",\"$4\",\"$5\",\"$6\",\"$8\",\"$9\",\"$10\",\"$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$28,$30'}"

func CollectProcessData() ([]string, error) {
	out, err := exec.Command("bash", "-c", ps_command).Output()
	if err != nil {
		return []string{}, err
	}
	return cleanProcessData(string(out)), nil
}

func cleanProcessData(process_data string) []string {
	data_slice := strings.Split(process_data, "\n")
	return data_slice
}
