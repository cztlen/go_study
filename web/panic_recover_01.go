package main

func main() {
	test()
}

//panic 抛出异常，recover接收异常
func test() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err.(string))
	// 	}
	// }()
	panic("panic")
}
