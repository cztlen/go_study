package main

import "fmt"

func main() {
	test()
}
func getCircle(radius float32) float32 {
	if radius < 0 {
		panic("半径不能为负")
	}
	return radius * radius
}
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()
	getCircle(-6) //若触发panic，不再运行
	fmt.Println("go on...")
}
