package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kartik1998/vegeta/scripts"
)

func main() {
	for t := range time.Tick(3 * time.Second) {
		_, err := scripts.CollectProcessData()
		if err != nil {
			log.Fatal(fmt.Sprintf("%s, %s", t, err))
			os.Exit(1)
		}
	}
}
