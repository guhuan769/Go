package main

import (
	"flag"
	"fmt"
	"time"
)

var j = flag.Int("j", 0, "")

func main() {
	flag.Parse()
	var i = 0
	for {
		fmt.Println("demo print", i, *j)
		i++
		time.Sleep(1 * time.Second)
	}
}
