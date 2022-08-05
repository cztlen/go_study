package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64          //共享变量
	wg      sync.WaitGroup //用来等待程序结束
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
		atomic.AddInt64(&counter, 1) //强制同一时刻只有一个goroutine运行并完成这个加法操作
		runtime.Gosched()            //goroutine 从线程退出，并放回到队列

	}
}
