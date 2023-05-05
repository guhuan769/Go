package main

import hello "HellowGolang/pakage"
import hello2 "HellowGolang/pakage/package2"

// 程序入口 func 声明
func main() {
	/*
		fmt.Println("Hello Golang")
		var str1 string
		var num int
		str1, num = f1("hellogolang f1")
		:= 给变量赋值得同时去声明变量
		string1, num1 := f1("hello golang string1 - 1 f1")
		fmt.Println(str1, num)
		fmt.Println(string1, num1)*/
	hello.F1()
	hello2.F2()
}
