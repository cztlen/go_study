package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("done\n")
				return
			case ch <- n:
				n++
				fmt.Printf("ch\n")
				time.Sleep(time.Second)

			}
		}
	}()
	return ch
}
func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
}
