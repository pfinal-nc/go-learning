/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-17 10:48:56
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 创建一个新的目录
	err := os.Mkdir("test", 0755)
	check(err)
	defer os.RemoveAll("test")

	// 创建一个用于创建临时文件的函数
	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}
	createEmptyFile("test/file1")
}
