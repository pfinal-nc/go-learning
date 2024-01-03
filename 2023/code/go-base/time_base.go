package main

import (
	"fmt"
	"time"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/1/3
 * @Desc:
 * @Project: 2023
 */

func main() {
	p := fmt.Println
	now := time.Now()
	_, _ = p(now)
	then := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	_, _ = p(then)
	_, _ = p(then.Year())
	_, _ = p(then.Month())
	_, _ = p(then.Before(now))
	// 时间戳
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano() / 1000000)
	fmt.Println(time.Unix(time.Now().Unix(), 0))
}
