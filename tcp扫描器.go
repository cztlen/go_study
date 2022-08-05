package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

//端口扫描器
func main() {
	host := flag.String("hostname", "baidu.com", "hostname to test")
	startPort := flag.Int("start port", 80, "the port on which the scanning starts")
	endPort := flag.Int("end port", 100, "the port on which the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond*200, "timeont")
	flag.Parse()
	ports := []int{}
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for port := *startPort; port < *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*host, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	fmt.Printf("opened ports:%v\n", ports)
}

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		conn.Close()
		fmt.Println("connect success ")
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
