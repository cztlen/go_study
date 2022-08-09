package main

import "fmt"

func main() {
	var mapSlice = make([]map[string]string, 2)
	for key, v := range mapSlice {
		fmt.Printf("index:%d value:%v\n", key, v)
	}
	fmt.Println("after init")
	mapSlice[0] = make(map[string]string, 2)
	mapSlice[0]["name"] = "peter"
	mapSlice[0]["age"] = "11"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
