/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-03 14:08:35
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
)

func f(form string) {
	for i := 0; i < 3; i++ {
		fmt.Println(form, ":", i)
	}
}

func main() {
	// 同步(sysnchronously) 调用
	f("direct")

	// 使用 `go f(s)` 在一个 Go 协程中 调用这个函数
	// 这个新的 Go 协程将会并发 执行这个函数
	go f("goroutine")

	// 你也可以匿名函数启动一个 go协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在这2个 Go 协程在独立的GO 协程中异步运行 所以在程序直接运行这行
	fmt.Scanln()
	fmt.Println("done")
}
