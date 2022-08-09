package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//这个只是一个简单的版本只是获取QQ邮箱并且没有进行封装操作，另外爬出来的数据也没有进行去重操作
var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
)

// 爬邮箱
func GetEmail() {
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "Http.get url")
	defer resp.Body.Close()
	// fmt.Println(resp.Body)
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	pageStr := string(pageBytes)
	// fmt.Println(pageStr)
	re := regexp.MustCompile(reQQEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	// fmt.Println(results)
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}

}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
func main() {
	GetEmail()
}
