package main

import (
	"fmt"
	"sync"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/6/24
 * @Desc:
 * @Project: 2024
 */

func main() {
	var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("你好世界")
			wg.Done()
		}()
	}
	// 等待N个后台线程完成
	wg.Wait()

}
