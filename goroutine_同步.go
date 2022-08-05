package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int //共享变量
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched() //goroutine 从线程退出，并放回到队列
		value++
		counter = value
	}
}
