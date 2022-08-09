package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var per = nePerson("我", 19)
	per.Dream()
	// fmt.Printf("%#v\n", per)
}
func nePerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
func (p *person) Dream() {
	fmt.Printf("%s 的梦想是学好go语言", p.name)
}
