package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name Name
	Age  int
}
type Name struct {
	First string
	Last  string
}

func main() {
	j := []byte(`{"name":{"first":"www.topgoer.com","last":"dj"},"age":18}`)
	var u User
	err := json.Unmarshal(j, &u)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("age: %d\n", u.Age)
	fmt.Printf("%s\n%s\n", u.Name.First, u.Name.Last)
}
