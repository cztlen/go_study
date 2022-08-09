package main

import (
	"context"
	"fmt"
	"time"
)

//取消控制
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Speak(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)
}

func Speak(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println("shutdown")
			return
		default:
			fmt.Println("bb")
		}

	}
}
