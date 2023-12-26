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

func workerC(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go workerC(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 9; a++ {
		<-results
	}
}

// 将多个并发实例中支持的任务了, 执行者将从 jobs 通道接收任务, 并且通过 results 发送对应的结果, 将让每个任务简介1s来模仿一个耗时的任务
// 为了使用 worker 工作池并且收集他们的结果, 需要2个通道. 启动了3个worker,初始是阻塞的, 因为还没有传递任务.这里发送9个jobs, 然后close 这些通道来表示这些就是
// 所有的任务
// 收集所有这些任务的返回值. 执行这个程序, 显示9个任务被多个worker执行.整个程序处理所有的任务仅执行3s而不是9s 是因为3个worker是并行的
