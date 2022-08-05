package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	failedMsg := "测试 in port:"
	writeLog(failedMsg, "./stat.log")
}

//写一个log生成文件
func writeLog(msg, logPaht string) {
	fd, _ := os.OpenFile(logPaht, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0664)
	defer fd.Close()
	// content := strings.Join([]string{msg, "\r\n"}, "ok") //字符串连接
	// buf := []byte(content)                               //将字符串转为字节切片
	// fmt.Println(content)
	// fd.Write(buf)
	var builder strings.Builder
	builder.WriteString("ACE")
	builder.WriteString("ACE")
	a := builder.String()
	fmt.Println([]byte(a))
	fd.Write([]byte(a))

}
