package main

import "fmt"

type fileError struct {
}

func (fe *fileError) Error() string {
	return "文件错误"
}
func main() {
	conent, err := openFile()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(conent))
	}
}
func openFile() ([]byte, error) {
	return nil, &fileError{}
}
