package main

import (
	"context"
	"fmt"
	"time"
)

func f1(ch1 chan struct{}) {
	time.Sleep(1 * time.Second)
	ch1 <- struct{}{}
}

func f2(ch2 chan struct{}) {
	time.Sleep(3 * time.Second)
	ch2 <- struct{}{}
}

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	go func() {
		go f1(ch1)
		select {
		case <-ch1:
			fmt.Println("f1 done")
		case <-ctx.Done():
			fmt.Println("time out")
			break
		}
	}()
	go func() {
		go f2(ch2)
		select {
		case <-ch2:
			fmt.Println("f2 done")
		case <-ctx.Done():
			fmt.Println(" f2 time out")
			break
		}
	}()
	time.Sleep(5 * time.Second)
}
