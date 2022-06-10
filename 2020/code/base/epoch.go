/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-15 11:53:04
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 分别使用带 Unix 或者 UnixNano 的time.Now()函数获取当前时间
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(secs)
	fmt.Println(nanos)

	// UnixMills 是不粗在的 所以要得到毫秒数 需要手动转化一下
	millis := nanos / 1000000
	fmt.Println(millis)

	// 你也可以将 Unix 纪元起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
