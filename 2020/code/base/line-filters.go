/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-16 11:54:02
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 对 os.Stdin  使用一个带缓冲的 scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Text() 返回当前的 token 输入的下一行
		ucl := strings.ToUpper(scanner.Text())
		// 打印行
		println(ucl)
	}
	// 检查 Scan 的错误  文件结束符是可以接受的
	// 检查 `Scan` 的错误。文件结束符是可以接受的，并且
	// 不会被 `Scan` 当作一个错误。
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
