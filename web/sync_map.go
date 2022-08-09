package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	scene.Store("bobo", 22)
	scene.Store("lailai", 29)
	scene.Store("cc", 31)
	scene.Delete("bobo")
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iter:", key, value)
		return true
	})
	fmt.Println(scene.Load("bobo"))

}
