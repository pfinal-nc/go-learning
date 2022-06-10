/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-03 17:17:51
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 这个 `rang` 迭代从 queue 中得到的每个值 因为我们在前面 close 了这个通道,这个迭代会在接收完2个值之后结束,如果我们没有close
	// 将在这个循环中阻塞执行, 等待接收第三个值
	for elem := range queue {
		fmt.Println(elem)
	}
}
