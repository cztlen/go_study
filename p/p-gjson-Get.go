package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	json := `{"name":{"first":"www.topgoer.com","last":"dj"},"age":18}`
	age := gjson.Get(json, "age")
	fmt.Println(age.Int())
	name := gjson.Get(json, "name")
	fmt.Println(name)
	first := gjson.Get(json, "name.first")
	fmt.Println(first.String())
}
