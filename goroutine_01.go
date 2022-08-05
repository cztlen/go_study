package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("CREATE GOROUTINES")
	go printss("a")
	go printss("b")
	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("terminating program")

}
func printss(s string) {
	defer wg.Done()
next:
	for count := 2; count < 5000; count++ {
		for i := 2; i < count; i++ {
			if count%i == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", s, count)
	}
	fmt.Println("complete", s)
}
