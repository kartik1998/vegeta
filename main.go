package main

import (
	"fmt"

	"github.com/kartik1998/freeza/scripts"
)

// ps aux | head -1; ps aux | sort -rnk 4 | head -5

func main() {
	data := scripts.CollectProcessData()
	fmt.Println(data)
}
