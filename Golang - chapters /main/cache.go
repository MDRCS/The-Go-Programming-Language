package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
} {
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}


func main() {
	var key = "Med"
	cache.mapping[key] = "AI SWE"

	value := Lookup(key)

	fmt.Print(value)
	//fmt.Println(cache)
}
