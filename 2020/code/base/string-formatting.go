/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-10 15:06:09
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import "fmt"

type proint struct {
	x, y int
}

func main() {

	p := proint{1, 2}
	//打印 point 结构体的实例
	fmt.Printf("%v\n", p)

	// 如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名
	fmt.Printf("%+v\n", p)

	// `%#v` 形式则输出这个值的 Go 语法表示
	fmt.Printf("%#v\n", p)

	// 需要打印纸的类型 使用 %T
	fmt.Printf("%T\n", p)
	// 格式化布尔值
	fmt.Printf("%t\n", true)
	// 格式化整形数
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 14)  //二进制
	fmt.Printf("%c\n", 33)  // 十进制转换为ASCII
	fmt.Printf("%x\n", 456) // 将数字转换为十六进制

	fmt.Printf("%f\n", 78.9)        // 浮点数
	fmt.Printf("%e\n", 123400000.0) // 科学计数法
	fmt.Printf("%E\n", 123400000.0) // 科学计数法

	fmt.Printf("|%d|%d|\n", 12, 345)
	fmt.Printf("|%6f|%6f|\n", 1.2, 3.45)
}
