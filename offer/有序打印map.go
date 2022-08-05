package main

import (
	"fmt"
	"sort"
)

func main() {
	// var a []string
	b := make(map[string]interface{})
	b["beijing"] = "首都"
	b["shanghi"] = "经济"
	b["xiaowang"] = "小王"
	b["xiaoming"] = "小c"
	sortPrint(b)
}

//有序打印map   sort.Ints  sort.Strings sort.Float
func sortPrint(m map[string]interface{}) {
	var s []string
	for k, _ := range m {
		s = append(s, k)
	}
	sort.Strings(s)

	for _, v := range s {
		fmt.Println(m[v])
	}

}

//
