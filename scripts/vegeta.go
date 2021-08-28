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

const ps_command = "ps aux | head -1 | awk {'print $1\",\"$2\",\"$3\",\"$4\",\"$5\",\"$6\",\"$8\",\"$9\",\"$10\",\"$11'}; ps aux | sort -rnk 4 | awk {'print $1\",\"$2\",\"$3\",\"$4\",\"$5\",\"$6\",\"$8\",\"$9\",\"$10\",\"$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$28,$30'}"

var log_file_name = utils.GetLogFileName()

func CollectProcessData() ([]string, error) {
	out, err := exec.Command("bash", "-c", ps_command).Output()
	if err != nil {
		return []string{}, err
	}
	data_slice := cleanProcessData(string(out))
	filepath, e := utils.CreateFile("./" + log_file_name)
	if e != nil {
		return []string{}, e
	}
	writeProcessData(filepath, data_slice)
	return data_slice, nil
}

func writeProcessData(filepath string, values []string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	for _, value := range values {
		fmt.Println(value)
		if _, err := file.WriteString(value + "\n"); err != nil {
			log.Fatal(err)
		}
	}
}

func cleanProcessData(process_data string) []string {
	data_slice := strings.Split(process_data, "\n")
	return data_slice
}
