package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/2/28 17:38
 * @Desc:
 */

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	user := User{Name: "PFinal", Email: "lampxiezi@163.com"}
	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println(result)
}
