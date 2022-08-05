package main

import (
	"fmt"
	"sort"
)

func main() {

	//字符串数组排序
	// myList := []string{"1", "10", "11", "2", "3", "4", "5", "6", "7", "8", "9"}

	// fmt.Printf("Before: %v\n", myList)

	// sort.Slice(myList, func(i, j int) bool {
	// 	numA, _ := strconv.Atoi(myList[i])
	// 	numB, _ := strconv.Atoi(myList[j])
	// 	return numA < numB
	// })
	//整型数组排序
	myInt := []int{1, 10, 11, 2}
	fmt.Printf("Before: %v\n", myInt)
	sort.Slice(myInt, func(i, j int) bool {
		return myInt[i] < myInt[j]
	})
	fmt.Printf("After: %v\n", myInt)
}
