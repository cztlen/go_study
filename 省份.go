package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

//一份数据，写到不同的省份文件
type Province struct {
	Id       string
	Name     string
	File     *os.File
	chanData chan string
}

var wg sync.WaitGroup

func main() {
	//声明map，存放所有省市
	pMap := make(map[string]*Province)
	ps := []string{
		"北京市11", "天津市12", "河北省13",
		"山西省14", "内蒙古自治区15", "辽宁省21", "吉林省22",
		"黑龙江省23", "上海市31", "江苏省32", "浙江省33", "安徽省34",
		"福建省35", "江西省36", "山东省37", "河南省41", "湖北省42",
		"湖南省43", "广东省44", "广西壮族自治区45", "海南省46",
		"重庆市50", "四川省51", "贵州省52", "云南省53", "西藏自治区54",
		"陕西省61", "甘肃省62", "青海省63", "宁夏回族自治区64", "新疆维吾尔自治区65",
		"香港特别行政区81", "澳门特别行政区82", "台湾省83"}

	for _, p := range ps {
		name := p[:len(p)-2]
		id := p[len(p)-2:]
		province := Province{Id: id, Name: name}
		pMap[id] = &province
		file, _ := os.OpenFile("./"+province.Name+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		province.File = file
		defer file.Close()
		province.chanData = make(chan string, 1024)
		fmt.Println(name, "管道已创建")

	}
	for _, province := range pMap {
		wg.Add(1)
		go writeFile(province)
	}
	file, _ := os.Open("./k.txt")
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			for _, province := range pMap {
				close(province.chanData)
				fmt.Println(province.Name, "管道已经关闭")
			}
			break
		}
		lineStr := string(lineBytes)
		fieldsSlice := strings.Split(lineStr, ",")
		id := fieldsSlice[1][0:2]
		if province, ok := pMap[id]; ok {
			province.chanData <- (lineStr + "\n")
		} else {
			fmt.Println("未知的省", id)
		}
	}
	wg.Wait()
}
func writeFile(province *Province) {
	defer wg.Done()
	for lineStr := range province.chanData {
		province.File.WriteString(lineStr)
		fmt.Println(province.Name, "写入", lineStr)
	}
}
