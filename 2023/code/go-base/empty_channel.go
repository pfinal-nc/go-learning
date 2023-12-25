package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/25
 * @Desc:
 * @Project: 2023
 */
func worker(ch chan struct{}) {
	<-ch
	fmt.Println("do something...")
	close(ch)
}

func main() {
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}
}

// 有时候使用channel 不需要发送任何的数据, 只用来通知子协程（goroutine）执行任务,或只用来控制协程并发度
