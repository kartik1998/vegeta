package main

import (
	"log"
	"os"
	"time"

	"github.com/kartik1998/freeza/scripts"
)

// ps aux | head -1; ps aux | sort -rnk 4 | head -5
// https://stackoverflow.com/questions/10781516/how-to-pipe-several-commands-in-go

func main() {
	for t := range time.Tick(3 * time.Second) {
		_, err := scripts.CollectProcessData()
		if err != nil {
			log.Fatal(t, err)
			os.Exit(1)
		}
	}
}
