package main

import (
	"fmt"

	"github.com/kartik1998/freeza/scripts"
)

// ps aux | head -1; ps aux | sort -rnk 4 | head -5
// https://stackoverflow.com/questions/10781516/how-to-pipe-several-commands-in-go

func main() {
	data := scripts.CollectProcessData()
	fmt.Println(data[0] + "\n" + data[1])
}
