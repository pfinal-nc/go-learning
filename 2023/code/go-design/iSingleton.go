package main

import "fmt"

// 单利模式采用了 饿汉式 和 蓝汉式 的实现

// 饿汉式

// Singleton 饿汉式 单利
type Singleton struct {
}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

// GetInstance 获取实例
func GetInstance() *Singleton {
	return singleton
}

func main() {
	fmt.Println(GetInstance() == GetInstance())
}
