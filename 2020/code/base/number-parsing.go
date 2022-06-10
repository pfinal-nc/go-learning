/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-15 16:07:43
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//使用 ParseFloat 解析浮点数 这里的 64 表示 解析的数的位数
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 在使用 ParseInt 解析整数时  0 表示自动推断字符串所表示的数字的进制
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi 是一个基础的10 进制整型数转换函数
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 在输入错误时候 解析函数会返回一个错误
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
