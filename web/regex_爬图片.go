package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// w代表大小写字母+数字+下划线
// reEmail = `\w+@\w+\.\w+`
// s?有或者没有s
// +代表出1次或多次
//\s\S各种字符
// +?代表贪婪模式
// reLinke  = `href="(https?://[\s\S]+?)"`
// rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
// reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
// reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`

//1.通过url获取图片资源
//2.将图片资源，url放入通道内
//3.将图片资源保存到本地
var (
	taskChan   chan string
	imagesChan chan string
	s          sync.WaitGroup
	reImg      = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

//url https://www.bizhizu.cn/shouji/tag-%E5%A5%B3%E7%A5%9E/8.html
func main() {
	taskChan = make(chan string, 26)
	imagesChan = make(chan string, 1000000)
	for i := 1; i < 2; i++ {
		s.Add(1)
		go getImages("https://www.bizhizu.cn/shouji/tag-%E5%A5%B3%E7%A5%9E/" + strconv.Itoa(i) + ".html")
	}

	s.Add(1)
	go checkOk()
	for i := 0; i < 20; i++ {
		s.Add(1)
		go down()
	}
	s.Wait()
	fmt.Println("success")
}

//1.通过url获取图片资源
func getImagesUrls(url string) (urls []string) {
	resp, err := http.Get(url)
	handFunc(err, "http.Get error")
	defer resp.Body.Close()

	pageBytes, err := ioutil.ReadAll(resp.Body)
	handFunc(err, "ioutil.ReadAll error")
	pageStr := string(pageBytes)

	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d张照片\n", len(results))
	for _, result := range results {
		urls = append(urls, result[0])
	}
	NAVH-2NET-33DX-5LJT
	return
}

//将url,图片资源放入对应的通道
func getImages(url string) {
	defer s.Done()
	urls := getImagesUrls(url)
	for _, url := range urls {
		imagesChan <- url
	}
	taskChan <- url
	return
}

//下载图片资源到本地
func down() {
	defer s.Done()
	for url := range imagesChan {
		filename := getFileNameFromUrl(url)
		res := downLoad(url, filename)
		if res != nil {
			fmt.Printf("%s 下载失败\n", filename)
		} else {
			fmt.Printf("%s 下载成功\n", filename)
		}
	}

}
func downLoad(url, filename string) error {
	resp, err := http.Get(url)
	handFunc(err, "http.Get error")
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	handFunc(err, "ioutil.ReadAll error")
	filename = "d:/images/" + filename
	err = ioutil.WriteFile(filename, bytes, 0666)
	return err

}

func getFileNameFromUrl(url string) (filename string) {
	lastIndex := strings.LastIndex(url, "/")
	//重命名
	temp := strconv.Itoa(int(time.Now().UnixNano()))
	filename = temp + "_" + url[lastIndex+1:]
	return

}

//检查任务通道是否完成任务
func checkOk() {
	defer s.Done()
	var count int
	for {
		url := <-taskChan
		fmt.Printf("%s 任务已完成\n", url)
		count++
		if count == 26 {
			close(imagesChan)
			break
		}
	}
	return
}

//处理异常
func handFunc(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
	return
}
