package main

import (
	"fmt"
	"time"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/12/28 22:43
 * @Desc:
 */

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second) // 模拟耗时操作
		done <- true
	}()

	for  {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-ticker.C:
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}