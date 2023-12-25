package main

import (
	"fmt"
	"time"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/25
 * @Desc:
 * @Project: 2023
 */

func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

// 1. 利用 time.After 启动一个异步的定时器, 返回一个 channel, 当超时指定的时间后, 该 channel将会接受到信号
// 2. 启动了子协程执行函数f, 函数执行结束后,将向 channel done 发送结束信号
// 3. 使用select阻塞等待 done 或 time.After 的信息, 若超时, 则返回错误. 若没有超时, 则返回 nil
func main() {
	err := timeout(doBadthing)
	if err != nil {
		return
	}
}
