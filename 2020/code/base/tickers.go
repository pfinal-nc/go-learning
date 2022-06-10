/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-03 17:55:58
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// ticker 定时器
	// 这里我们在这个通道上使用内置的 `range` 来迭代值
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// 打点器可以和定时器一样被停止  一旦一个打点停止了, 将不能再从它的通道中接收到值
	// 将在运行后1600ms 停止这个打点器
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
