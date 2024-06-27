package main

import (
	"fmt"
	"os"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/6/27
 * @Desc:
 * @Project: 2024
 */

func main() {
	defer fmt.Printf("!") // 使用 os.Exit(3) 时 defer 将不会执行, 所以这里的 fmt.Println将永远不会被调用 退出并且退出状态为3
	os.Exit(3)            // 如果使用 go run 来 运行 exit.go 那么退出状态将会被go 捕获并打印
	// 使用编译并执行一个二进制文件的方式退出
}
