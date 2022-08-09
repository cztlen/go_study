package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool, 1)
	go func(temp chan bool) {
		time.Sleep(time.Second * 5)
		temp <- true
	}(ch)
	// ch <- true
	select {
	case <-ch:
		{
			fmt.Println("ch")
		}
	case <-time.After(time.Second * 10):
		fmt.Println("10")
		// default:
		// 	fmt.Println("default")
	}
}
