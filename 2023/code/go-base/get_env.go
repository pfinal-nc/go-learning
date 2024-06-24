package main

import (
	"fmt"
	"os"
	"strings"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/1/5
 * @Desc:
 * @Project: 2023
 */

func main() {
	//_ = os.Setenv("FOO", "1")
	//fmt.Println(os.Getenv("FOO"))
	fmt.Println(os.Getenv("BAR"))
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0] + "=" + pair[1])
	}
}
