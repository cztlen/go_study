package main

import (
	"fmt"
	"time"
)

//仅包含方法的结构体

func main() {
	fmt.Println(timeout(doBad))

}
func doBad(ch chan bool) {
	time.Sleep(time.Second)
	ch <- true
}
func timeout(f func(chan bool)) error {
	ch := make(chan bool)
	go f(ch)
	select {
	case <-ch:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		// fmt.Println("t")
		return fmt.Errorf("TIMEOUT")
	}

}
