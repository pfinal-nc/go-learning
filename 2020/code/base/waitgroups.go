/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-06 15:56:24
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个协程都会运行该函数 WaitGroup 必须通过指针传递给函数
func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("worker %d starting\n", id)

	// 睡眠一秒钟  以此来模拟耗时的任务
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)

	//通知 WaitGroup 当前协程的工作已完成
	wg.Done()
}

func main() {

	// 这个 WaitGroup 被用于 等待该函数开启的所有的协程。
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// 阻塞, 知道 WaitGroup 计数器恢复为0 既所有的协程工作都完成
	wg.Wait()
}
