package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int
}

func main() {
	s1 := Student{
		Name: "小米",
		Age:  5,
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("marshal failed")
		return
	}
	fmt.Printf("%s\n", data)
	c1 := &Student{}
	err = json.Unmarshal([]byte(data), c1)
	if err != nil {
		fmt.Println("unmarshal faild")
		return
	}
	fmt.Printf("%v\n", c1)
}
