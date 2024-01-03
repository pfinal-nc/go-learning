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
	p(then.Year())
	p(then.Month())
}
