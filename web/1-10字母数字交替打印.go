package main

import (
	"fmt"
	"time"
)

var (
	numCh  = make(chan struct{}, 1)
	charCh = make(chan struct{}, 1)
)

func printNum() {
	for i := 0; i < 10; i++ {
		<-numCh
		fmt.Println(i)
		charCh <- struct{}{}
	}
}
func printChar() {
	for i := 'a'; i < 'j'; i++ {
		<-charCh
		fmt.Println(string(i))
		numCh <- struct{}{}
	}
}

func main() {
	numCh <- struct{}{}
	go printNum()
	go printChar()
	time.Sleep(time.Second)
}
