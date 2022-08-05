package main

import (
	"fmt"
)

type notifier interface {
	notify()
}
type user struct {
	Name string
	Age  int
}

func (u *user) notify() {
	fmt.Println(u.Name, u.Age)
}
func main() {
	u := &user{
		Name: "小米",
		Age:  4,
	}
	sendNotify(u)
}

func sendNotify(n notifier) {
	n.notify()
}
