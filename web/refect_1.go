package main

import (
	"fmt"
	"reflect"
)

//字符串函数名调用函数
type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("11")
}
func main() {
	a := Animal{}
	reflect.ValueOf(&a).MethodByName("Eat").Call([]reflect.Value{})
}
