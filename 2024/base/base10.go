package main

import (
	"fmt"
	"time"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/6/27
 * @Desc:
 * @Project: 2024
 */
// 超时 对于一个连接外部资源, 或者其它一些需要花费执行时间的操作的程序而言是很重要的.得益于通道和 select

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2) // 假如我们执行一个 外部调用, 并在2秒后通过通道c1 返回它的执行结果
		c1 <- "result"
	}()
	select { // 这里是使用 select 实现一个超时操作. res := <- c1 等待结果, <- Time.After 等待超时时间1秒后发送的值. 由于 select 默认处理第一个已准备好的接收操作,
	// 如果这个操作超过了允许的 1秒的话,将会执行超时 case. 如果我允许一个长一点的超时时间3秒, 将会成功的从c2接收到值, 并且打印出结果
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result"
	}()
	select { // 运行这个程序, 首先显示运行超时的操作, 然后是成功接收的.
	//使用这个 select 超时方式, 需要使用通道传递结果。这对于一般情况是个好的方式,因为其他重要的 go 特性是基于通道和seleect的
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
