package main

import (
	"sync"
)

// ćšć±ćé
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
