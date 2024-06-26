package main

import (
	"fmt"
	"runtime"
	"sync"
)

//
//func say(s string) {
//	for i := 0; i < 3; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//}
//
//func main() {
//	go say("hello world")
//	time.Sleep(1000 * time.Millisecond)
//	fmt.Println("over!")
//}

// goroutine 生命周期

func main() {
	//  分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU())
	// wg 用来等待程序完成
	// 计数加2 表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数 并创建一个goroutines
	go func() {
		// 在函数退出调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()
		// 显示字符表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()
	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	// 等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}
