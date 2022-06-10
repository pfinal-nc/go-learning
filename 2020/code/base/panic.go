/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-08-30 16:08:15
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import "os"

func main() {
	panic("a program")

	// panic 的一个基本用法就是在函数返回了错误值 但是不知道或者不想处理时候终止
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
