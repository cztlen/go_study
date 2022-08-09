package main

import (
	"context"
	"fmt"
	"time"
)

func NewContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}
func HandlerFunc() {
	ctx, cancle := NewContextWithTimeout()
	defer cancle()
	deal(ctx)

}

//达到超时时间中止接下来的执行
func deal(ctx context.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
		}

	}
}
func main() {
	HandlerFunc()
}
