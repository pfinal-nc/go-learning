package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/25
 * @Desc:
 * @Project: 2023
 */
// 当使用通道作为函数的参数时, 可以指定通道是不是只用来发送或者接收值.
func ping(pings chan<- string, msg string) {
	// ping 函数定义一个只允许发送数据的通道, 尝试使用这个通道来接收数据将会得到一个编译时错误.
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// pong 函数允许通道来接收数据,另一个通道 pongs 来发送数据
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passwd message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
