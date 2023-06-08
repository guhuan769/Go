package main

import (
	"fmt"
	"memCache/cache"
)

const (
	B = 1 << (iota * 10) // iota 是第0个值 0*10 等于0 左移0位还是1
	kb
	MB
	GB
	TB
	PB
)

func main() {
	//c := cache.NewMemCache()
	//c.Get("")
	//Exec(c)
	//fmt.Println(B, kb, MB)
	fmt.Println(cache.ParseSize("1B"))
	//fmt.Println(cache.ParseSize("1kb"))
	//fmt.Println(cache.ParseSize("1mb"))
	//fmt.Println(cache.ParseSize("1gb"))
	//fmt.Println(cache.ParseSize("1tb"))
	//fmt.Println(cache.ParseSize("1tbGG"))
	cache.GetValSize(1)
	cache.GetValSize("sdasdasdsadsadsadsadsads")
}

func Exec(c cache.Cache) {

}
