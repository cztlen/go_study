package main

import "fmt"

//延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
func main() {
	test()
}
func test() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}
