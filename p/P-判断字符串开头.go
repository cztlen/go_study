package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "sex is moginc"
	ser := "Q"
	bool := strings.HasPrefix(str, ser)
	if bool {
		fmt.Println("true")
	} else {
		fmt.Println("FALSE")
	}

}
