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
// 定时器是需要再未来某一刻执行一次时使用的
// 打点器则是当想要在固定的时间间隔重复执行准备的 打点器和定时器的机制有点相似: 一个通道用来发送数据.在这个通道上使用内置的range 来迭代值每个 500ms 发送一次的值

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 16000)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
