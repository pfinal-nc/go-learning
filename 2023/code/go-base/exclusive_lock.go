package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/26
 * @Desc:
 * @Project: 2023
 */

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops int64 = 0
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
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
	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

// mutex 将同步对 state的访问, 每次循环读取, 使用一个键来进行访问, Lock() 这个 mutex 来确保对 state的独占访问, 读取选定的键的值, Unlock() 这个mutex 并且 ops值加1
// 为了确保这个 Go协程不会再调度中饿死, 在每次操作后明确的使用 runtime.Gosched()进行释放, 这个释放一般是自动处理的. 像例如每个通道操作后或者 time.Sleep 的阻塞调用后相似, 但是在
// 这个例子中需要手动的处理。同样的, 运行 10个Go协程来模拟写入操作, 使用和读取相同的模式
