package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/6/27
 * @Desc:
 * @Project: 2024
 */

func main() {
	sigs := make(chan os.Signal, 1) // go 通过向一个通道发送 os.Signal 值来进行信号通知, 将创建一个通道来接收这些通知
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // signal.Notify 注册这个给定的通道用于接收特定的信号.
	go func() {                                          // 这个go 协程执行一个阻塞的信号接收操作, 当它得到一个值时, 他将打印这个值然后通知程序可以退出
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done // 程序将在这里进行等待,直到它的到了期望的信号, 也就是上面的 go 协程发送的 done值 然后退出
	fmt.Println("exiting")
	// 当我们运行这个程序时, 它将一直等待一个信号, 使用 ctrl-C 终端显示未c 可以发送一个SIGINT 信号, 这会使程序打印interrrupt 然后退出
}
