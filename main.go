package main

import (
	"log"
	"os"
	"time"

	"github.com/kartik1998/freeza/scripts"
)

func main() {
	for t := range time.Tick(3 * time.Second) {
		_, err := scripts.CollectProcessData()
		if err != nil {
			log.Fatal(t, err)
			os.Exit(1)
		}
	}
}
