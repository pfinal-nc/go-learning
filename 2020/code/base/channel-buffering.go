/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-03 14:20:03
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
)

func main() {
	// 这里我们创建一个字符串通道 最多允许缓存2个值
	message := make(chan string, 2)

	// 由于此通道是缓冲的, 因此我们可以将这些值发送到通道中
	message <- "pfinal"
	message <- "南丞"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
