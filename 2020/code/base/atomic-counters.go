/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-06 16:49:35
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	// 使用一个无符号整形表示 计数器
	var ops uint64 = 0

	// 启动 50个Go协程 对计数器每个 1ms 进行加一操作
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 使用 addUnit64 来让计数器自动增加 使用 &ops 语言来给出 ops 的内存地址
				atomic.AddUint64(&ops, 1)
				// 允许其他 Go协程进行执行
				runtime.Gosched()
			}
		}()
	}

	// 等待1秒 让ops的自加操作执行一会
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)

}
