package main

import (
	"fmt"
)

//向已关闭的通道发送数据导致panic
func main() {
	test()
}
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ch := make(chan int)
	close(ch)
	ch <- 1
}
