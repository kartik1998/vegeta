package scripts

import (
	"bytes"
	"log"
	"os/exec"
)

func CollectProcessData() string {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}
