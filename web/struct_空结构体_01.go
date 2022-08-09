package main

//使用channel不发送任何的数据，通知子协程执行任务，或只用来控制协程并发度
//使用空结构体只作为占位符，不占用内存
import (
	"fmt"
	"sync"
)

var s sync.WaitGroup

func worker(ch chan struct{}) {
	defer s.Done()
	<-ch
	fmt.Println("do something")
	close(ch)
}

func main() {
	s.Add(1)
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}
	s.Wait()

}
