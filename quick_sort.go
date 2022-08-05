package main

import "fmt"

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
}
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	split := arr[0]
	low := make([]int, 0, 0)
	hight := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, split)
	for i := 1; i < len(arr); i++ {
		if arr[i] > split {
			hight = append(hight, arr[i])
		} else if arr[i] < split {
			low = append(low, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)
	sort_arr := append(append(low, mid...), hight...)
	return sort_arr
}
