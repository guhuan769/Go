package main

import "memCache/cache"

func main() {
	c := cache.NewMemCache()
	c.Get("")
	Exec(c)
}

func Exec(c cache.Cache) {

}
