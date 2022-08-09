package main

import (
	"fmt"
	"sync"
)

var (
	ch    = make(chan struct{})
	count = 11
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	go odd()
	go even()
	wg.Wait()

}

//打印奇数
func odd() {
	defer wg.Done()
	for i := 1; i < count; i++ {
		ch <- struct{}{}
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

//打印偶数
func even() {
	defer wg.Done()
	for i := 1; i < count; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
