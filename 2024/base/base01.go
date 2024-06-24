package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/6/24
 * @Desc:
 * @Project: 2024
 */

func main() {
	// 创建一个无缓存的 channel
	ch := make(chan int) // 第二个参数为0 或者不写第二个参数都可以

	// len(ch) 缓冲区剩余的数据个数 cap(ch) 缓冲区大小,两者这里永远都是0
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))
	// 新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("准备写入第 %d 个数据\n", i)
			ch <- i
			fmt.Printf("写入第 %d 个数据\n", i)
		}
	}()
}
