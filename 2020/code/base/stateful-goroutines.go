/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-07 10:23:04
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 在这个例子中, state 将被 一个单独的  go 协程拥有,这就能够保证数据在并行读取时不会混乱, 为了对
// state 进行读取或写入 其他的Go协程 将发送一条数据到 channel, 并且在接收到数据之后, 将数据写入 state

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// 和前面一样  我们将计算执行操作的次数
	var readOps uint64 = 0
	var writeOps uint64 = 0

	// `reads` 和 `wirtes` 通道分别将被其他的Go 协程用来发布和写入
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 这个就是拥有 state 的那个 Go 协程  和前面的例子中map 一样 不过这里是被这个状态协程私有的
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动 100 个Go协程, 通过 reads 通道发对 state 所有者 Go协程的读取请求
	// 每个读取请求需要构造一个 readOp 发送它到  reads通道中 并通过给定的 resp 通道接收
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 用相同的方法 启动 10个 写操作
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 让Go 协程们跑 1s
	time.Sleep(time.Second)
	// 最后 获取 ops值
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println(readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println(writeOpsFinal)
}
