package main

import (
	"fmt"
	"sync"
)

var (
	count = 10
	wg    sync.WaitGroup
	ch    = make(chan struct{})
)

func main() {
	wg.Add(2)
	go oddNumPrint(ch)
	go evenPrint(ch)
	wg.Wait()
	// time.Sleep(time.Second)
}

//打印奇数
func oddNumPrint(ch chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		ch <- struct{}{}
		if i%2 == 1 {
			fmt.Println(i)

		}
	}
}

//打印偶数
func evenPrint(ch chan struct{}) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Println(i)

		}
	}
}
