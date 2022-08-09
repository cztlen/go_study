package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path       string
	op         string
	createTime string
	mes        string
}

func main() {
	err := test("/Users/5lmh/Desktop/go/src/test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:

	}
}

//pathError实现了Error方法，error接口类型的实例
func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path, p.op, p.createTime, p.mes)
}
func test(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return &PathError{
			path:       fileName,
			op:         "read",
			mes:        err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	defer file.Close()
	return nil
}
