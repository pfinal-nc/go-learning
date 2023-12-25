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

func workerS(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 2)
	go workerS(done)
	<-done
	fmt.Println("main done")
	// 将要在go协程中运行的函数,done通道将被用于通知其他 go 协程这个函数已经工作完毕
}
