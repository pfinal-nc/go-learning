/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-03 17:26:36
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	// 定时器表示在未来某一个时刻的独立事件
	timer1 := time.NewTimer(time.Second * 2)

	// `<- timer1.C` 直到这个定时器的通道 `C` 明确的发送了一个值 定时器失效的值之前,将一直阻塞
	<-timer1.C
	fmt.Println("Time 1 expired")

	// 如果你需要的仅仅是单纯的等待, 你需要史依弘 time.sleep
	// 定时器有用的原因之一就是你可以在定时器失效之前 取消定时器
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("TImer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("TImer 2 stopped")
	}
}
