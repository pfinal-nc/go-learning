/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-03 15:37:53
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import "fmt"

func main() {

	message := make(chan string)
	//	signal  := make(chan bool)

	// 在message 中 存在 然后select 将这个值带入 <-message case 中 如果不是  就直接到 default 分支中

	// 一个非 阻塞发送的实现方法和上面一样
	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("send message", <-message)
		fmt.Println("send message", <-message)
	default:
		fmt.Println("no message send")
	}

	select {
	case msg := <-message:
		fmt.Println("receive message", msg)
	default:
		fmt.Println("no message received")
	}
}
