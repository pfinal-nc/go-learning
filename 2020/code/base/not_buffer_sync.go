/*
 * @Author: PFinal
 * @Date: 2021-08-30 17:59:02
 * @LastEditTime: 2021-08-30 18:11:37
 */
package main 

import (
    "fmt"
    "time"
)

func main() {
	c := make(chan int)
    done := make(chan string)
	go func(ch chan<- int,x int) {
		time.Sleep(time.Second * 5)
		ch <- x*x*x
	}(c,10)	
	go func(ch <-chan int) {
		data := <-ch
		fmt.Println(data)
		s := "程序结束"
		done <- s
	}(c)
	fmt.Println("程序开始")
	fmt.Println(<-done)

}