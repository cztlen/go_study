package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/axgle/mahonia"
)

func main() {
	read()
}
func read() {
	file, _ := os.Open("E:/gamelog/logbus/30001/log.20220606.20220606.action.1801.30001.6505901.track.user_battle_multi.log")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		gbkStr := string(lineBytes)
		utfStr := ConvertEncoding(gbkStr, "GBK")
		fmt.Println(utfStr)
	}
}
func ConvertEncoding(srcStr string, encoding string) (dstStr string) {
	enc := mahonia.NewDecoder(encoding)
	dstStr = enc.ConvertString(srcStr)
	return
}
