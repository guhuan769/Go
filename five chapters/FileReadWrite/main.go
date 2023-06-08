package main

import (
	"log"
	"os"
)

func main() {
	log.Println("log log")
	// 0644 0 代表的不是目录 6代表可读可写
	f, _ := os.OpenFile("destFile/log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetOutput(f)
}
