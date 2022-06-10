/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-03 16:06:33
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import "fmt"

// 使用一个 jobs 通道来传递 main() 中 GO 协程任务执行结果到一个工作Go协程 中
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 这里使用 `jobs` 发送 3 个任务到工作函数中，然后
	// 关闭 `jobs`。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send jobs", j)
	}

	close(jobs)

	<-done
}
