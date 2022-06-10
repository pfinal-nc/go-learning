/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-06-15 10:55:24
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import "fmt"

// 通过两个函数:`zeroval`和 `zeroptr` 来比较指针和值类型的不同
// `zeroval` 有一个 `int`的参数
// `zeroval` 将从调用它的那个函数中得到一个 `ival` 形式的拷贝
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` 有一个 `*int` 参数意味着他接收一个整形的 指针

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println(i)

	// 通过 `&i` 语法来取得 `i` 的内存地址，即指向 `i` 的指针。
	zeroptr(&i)
	fmt.Println(i)
	// 指针也能够被打印
	fmt.Println(&i)
	// integer := make([]int, 20)
	// integer = append(integer, -5900)
	// for i := 0; i < len(integer); i++ {
	// 	fmt.Println(i)
	// }
	// s2 := integer[1:3]
	// s2[0] = -100
	// s2[1] = 123456
	// fmt.Println(s2)
}
