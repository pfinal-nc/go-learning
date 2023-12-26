package main

import (
	"runtime"
	"sync/atomic"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/26
 * @Desc:
 * @Project: 2023
 */

// go 中最主要的状态管理方式是通过通道间的沟通来完成的, 在工作池的例子中碰到过, 但是还是有一些其他的方法来管理状态
func main() {
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

}
