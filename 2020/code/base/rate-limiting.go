/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-06 16:12:03
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 限制接收请求的处理,
	// requests := make(chan int, 5)
	// for i := 1; i <= 5; i++ {
	// 	requests <- i
	// }
	// close(requests)

	// 速率限制任务重的管理器
	// limiter := time.Tick(time.Millisecond * 200)

	// // 通过在每次请求前阻塞 limiter 通道的一个接收 限制每次200msg 执行依次
	// for req := range requests {
	// 	<-limiter
	// 	fmt.Println("request", req, time.Now())
	// }
	// 有时候我们想临时进行速率限制 并且不影响整体的速率控制
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每 200 ms 我们将添加一个新的值到 `burstyLimiter`中
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()
	// 现在模拟超过 5 个的接入请求, 它们中刚开始 3个将  接收 burstyLimiter 的脉冲影响 可以快速执行
	burstyRequests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
