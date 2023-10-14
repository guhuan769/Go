package main

import "fmt"

func f1(str string) (string, int) {
	fmt.Println(str)
	return "f1 return" + str, 0
}
