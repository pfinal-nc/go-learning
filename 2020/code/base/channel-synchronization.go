/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-03 14:32:37
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值通知我们已经工作完工
	done <- true
}

func main() {
	// 运行一个 worker go 协程 并基于用于通知的通道

	done := make(chan bool, 1)
	go worker(done)

	// 程序将在接收到通道中的 worker 发出的通知前一直阻塞
	<-done
}
