package main

import (
	"errors"
	"fmt"
	"os"
)

const perm = 0777

//双向链表  非线性安全
func main() {
	path := "D:/go/mm"
	stat, _ := os.Stat(path)
	fmt.Println(stat)
	ok, _ := pahtExist(path)
	if ok {
		fmt.Println("文件夹已存在")
		return
	}

	os.Mkdir(path, perm)
	fmt.Println(path, "创建成功")

}
func pahtExist(path string) (bool, error) {
	stat, err := os.Stat(path)
	fmt.Printf("%v", stat)
	// return false,nil
	if err == nil {
		if stat.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
