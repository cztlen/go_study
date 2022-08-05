package main

import (
	"sync"
)

// 全局变量
var (
	m       sync.RWMutex
	counter int64
)

func main() {

	for i := 0; i < 1000; i++ {
		m.Lock()
		counter++
		m.Unlock()
	}

	println(counter)
}
