package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/25
 * @Desc:
 * @Project: 2023
 */

func main() {
	message := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("nomessage received")
	}

	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
