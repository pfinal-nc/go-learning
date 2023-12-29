package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/28
 * @Desc:
 * @Project: 2023
 */
type point struct {
	x, y int
}

// %#v 输出这个值的 go 语法表示
// %T 打印值的类型
// %d 使用十进制格式进行格式化
func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 14)
}
