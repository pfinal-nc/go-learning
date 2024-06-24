package main

import (
	"fmt"
	"os/exec"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/1/5
 * @Desc:
 * @Project: 2023
 */
// 生成进程
func main() {
	// 从一个简单的命令开始, 没有参数或者输入,仅打印一些信息到标准输出流 exec.Command 函数帮助创建一个表示这个外部进程的对象
	// .Output 是另一个帮助处理运行一个命令的常见情况的函数,他等待命令运行完毕
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date" + string(dateOut))
}
