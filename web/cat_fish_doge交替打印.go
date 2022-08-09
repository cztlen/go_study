package main

import (
	"fmt"
	"time"
)

var dog = make(chan struct{})
var cat = make(chan struct{})
var fish = make(chan struct{})

func Dog() {
	for i := 0; i < 18; i++ {
		<-dog
		fmt.Println("dog")
		cat <- struct{}{}
	}

}
func Cat() {
	for i := 0; i < 18; i++ {
		<-cat
		fmt.Println("cat")
		fish <- struct{}{}
	}

}

func Fish() {
	for i := 0; i < 18; i++ {
		<-fish
		fmt.Println("fish")
		dog <- struct{}{}
	}

}

func main() {
	go Dog()
	go Cat()
	go Fish()
	cat <- struct{}{}

	time.Sleep(1 * time.Second)
}
