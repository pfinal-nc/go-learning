/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-06 13:59:36
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"time"
)

// 这里是 worker 我们将并发执行多个 worker
// worker 将从 jobs 通道接收任务 并且 通过 resultes 发送对应的结果
// 我们将让每个任务间隔 1s 来模仿一个耗时

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	// 为了使用 worker 线程池并且收集他们的结果, 我们需要2个通道
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 这里启动3 个 worker 初始是阻塞的 因为没有传递任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 这里我们 发送 9个 `jobs` 然后 `close` 这些管道
	// 来表示这些就是所有的任务了
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	// 最后，我们收集所有这些任务的返回值。
	for a := 1; a <= 9; a++ {
		<-results
	}
}
