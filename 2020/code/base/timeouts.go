/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-03 14:55:37
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// 这里使用 select 实现一个超时操作, res := <-c1 等待结果, <-Time.After 等待超时时间1秒后发送的值
	// 由于 select 默认处理第一个
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout", <-c1)
	}

	// 如果允许一个长一点的超时时间 3秒 将会 成功的从 c2 接收到值 并且打印出来
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout ", <-c2)
	}
}
