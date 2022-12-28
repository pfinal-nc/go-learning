package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/12/28 22:34
 * @Desc:
 */

func main() {
	var number uint32
	var wg sync.WaitGroup

	for i := 0; i <1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint32(&number,1) // 使用原子操作
		}()
	}
	wg.Wait()
	fmt.Printf("number: %d\n", number)

}