package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/1/4
 * @Desc:
 * @Project: 2023
 */

// 一个行过滤器在读取标准输入流的输入, 处理该输入, 然后将得到一些的结果输出到标准输出的程序中常见的一个功能, grep 和 sed 是常见的一个功能.
// grep 和 sed 是常见的行过滤器

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}
