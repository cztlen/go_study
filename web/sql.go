package main

import (
	"database/sql"
	"fmt"
)

func main() {
	_, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	if err != nil {
		fmt.Println("bb", err)
	}
	fmt.Println("aa")
}
