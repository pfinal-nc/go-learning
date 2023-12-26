package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/26
 * @Desc:
 * @Project: 2023
 */

// 关闭一个通道意味着不能再想这个挺大发送只了, 这个特性可以用来给这个通道的接收方传达工作已经完成的信息
// 使用一个jobs 通道来传递 main中 go协程任务执行的信息到一个工作Go协程中,没有多余的任务给这个Go协程时, 将close这个jobs 通道
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
	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("set job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
