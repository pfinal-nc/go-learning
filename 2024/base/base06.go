package main

import (
	"fmt"
	"os/exec"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/6/26
 * @Desc:
 * @Project: 2024
 */
// 有时候, go 程序需要生成其他的, 非go进程. 这个网站的语法高亮是通过在 Go程序中生成一个 pygmentize 来实现的。
// 将从一个简单的命令开始,没有参数或者输入, 仅仅打印一些信息到标准输出流, exec.Command 函数 帮助创建一个表示这个外部进程的对象
// .Output 函数 是另一个帮助我们处理运行一个命令的常见情况的函数,它等待命令运行完成,并收集命令的输出, 如果没有出错, dateOut 将获取到日期信息的字节
//
func main() {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	// grepOut, _ := grepCmd.StdoutPipe()
	_ = grepCmd.Start()
	_, _ = grepIn.Write([]byte("hello grep\ngoodbye grep"))
	_ = grepIn.Close()
	_, _ = grepIn.Write([]byte("hello grep\ngoodbye grep"))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
