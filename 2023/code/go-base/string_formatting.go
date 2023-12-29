package main

import (
	"fmt"
	"os"
)

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
// %b 使用二进制格式进行格式化

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 33)
	fmt.Printf("%x\n", 456)                   // %x 使用十六进制编码
	fmt.Printf("%f\n", 78.9)                  // %f 进行最近本的十进制格式化
	fmt.Printf("%e\n", 12300000.0)            // %e 和%E将浮点型格式化为科学计数法表示形式
	fmt.Printf("%s\n", "\"string\"")          // %s 用于输出字符串
	fmt.Printf("%q\n", "\"string\"")          // %q 用于输出带双引号的字符串
	fmt.Printf("%x\n", "hex this")            // %x 用十六进制编码输出字符串
	fmt.Printf("%p\n", &p)                    // %p 输出一个指针
	fmt.Printf("|%6d|%6d|\n", 12, 345)        // 当输出数字的时候 经常想要控制输出结果的宽度和精度可以使用%后面使用数字来控制宽度
	fmt.Printf("|%-6.2f|%6.2f|\n", 1.2, 3.45) // 左右对齐使用 -
	fmt.Printf("|%6s|%6s|\n", "foo", "b")     // 当输出字符串的时候 想要控制输出字符的个数
	fmt.Printf("|%-6s|%6s|\n", "foo", "b")    // 当输出字符串的时候
	s := fmt.Sprintf("a %s", "string")        // Sprintf 格式化并返回一个字符串而不打印出来
	fmt.Println(s)
	_, _ = fmt.Fprintln(os.Stderr, "an error")
}
