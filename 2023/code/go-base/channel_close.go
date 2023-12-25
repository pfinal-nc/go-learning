package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/25
 * @Desc:
 * @Project: 2023
 */

// 默认通道是无缓冲的 这意味着只有在对应的接收, 通道准备好接收时, 才允许进行发送(chan <-) 可缓存通道允许在没有对应接收方的情况下,缓存限定数量的值

func main() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
