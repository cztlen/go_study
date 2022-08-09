package main

//仅包含方法的结构体
import (
	"fmt"
)

type Door struct{}

func (d Door) worker() {
	fmt.Println("let us play")
}

func main() {
	var d Door
	d.worker()
}
