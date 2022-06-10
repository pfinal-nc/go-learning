/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-06 17:08:31
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// 在我们的例子中 `state` 是 一个 map
	var state = make(map[int]int)

	// 这里的 mutex 将同步对 state 的访问
	var mutex = &sync.Mutex{}

	// 为了比较基于互斥锁的处理方式和我们后面将要看到的其他
	// 方式 ops 将记录 我们对 state 的操作次数
	var ops int64 = 0

	// 这里我们运行100个 Go 协程, 每个协程都会读取 state的值
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// 每次循环读取
				// Lock() 这个  mutex 来确保 state的独占访问, 读取选定的键  Unlock() 这个mutex 并且 ops 值加1
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				// 为了确保这个Go 协程不会一直运行，我们每次操作后明确使用 runtime.Gosched() 来让出CPU时间片
				// 这个释放一般是自动处理的例如每个通道操作后 或者 time.Sleep 的阻塞调用后相似
				runtime.Gosched()
			}
		}()
	}

	// 同样的  我们运行10个 go 协程模拟写入操作
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	// 让这 10 个 Go 协程对 `state` 和 `mutex` 的操作
	// 运行 1 s。
	time.Sleep(time.Second)

	// 获取并输出最终的操作计数。
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	// 对 `state` 使用一个最终的锁，显示它是如何结束的。
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
