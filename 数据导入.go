package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

type Account struct {
	Id         uint   `gorm:"AUTO_INCREMENT"`
	Account_id string `gorm:"size:100"`
	Event_time uint
}

func main() {
	Db, err := gorm.Open("mysql", "root:root@(localhost:3306)/wash?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Errorf("创建数据库连接失败:%v", err))

	}
	defer Db.Close()
	Db.AutoMigrate(Account{}) //在库里创建表

	file, err := os.Open("D:\\go\\ww.txt")
	if err != nil {
		fmt.Println(errors.New("文件读取异常"))
	}
	defer file.Close()
	read := bufio.NewReader(file) //new bufio 存储文件信息
	tmp := "2006-01-02 15:04:05"
	for {
		lineBytes, _, err := read.ReadLine()
		if err == io.EOF {
			fmt.Println("文件数据已全部读取")
			break
		}
		lineStr := string(lineBytes)

		lineSlice := strings.Split(lineStr, ",")

		timeStr := lineSlice[1]

		res, _ := time.ParseInLocation(tmp, timeStr, time.Local)

		a := Account{
			Account_id: lineSlice[2],
			Event_time: uint(res.Unix()),
		}

		Db.Create(&a)
		// break

	}
}
