package main

import (
	"fmt"
	"time"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/26
 * @Desc:
 * @Project: 2023
 */

// 定时器表示在未来某一时刻的独立事件, 高度定时器需要等待的时间, 然后它将提供一个用于通知的通道.
// <-timer1.C 直到这个定时器的通道C 明确的发送定时器失效的值之前, 将一直阻塞 如果需要的仅仅是单纯的等待, 需要使用 time.Sleep 定时器是有用原因之一就是在定时器失效之前
// 取消这个定时器

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	// 第一个定时器将在程序开始后 ~2s失效, 第二个在它没失效之前就停止了
}
