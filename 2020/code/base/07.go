/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-06-15 10:19:47
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import "fmt"

// 闭包函数

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	fmt.Println(intSeq()) // 返回函数的地址

	nextInt := intSeq()

	// 通过多次调用来看闭包函数的效果
	fmt.Println(nextInt(), nextInt())

	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}
		if i == 95 {
			break
		}
		fmt.Print(i, "")
		fmt.Println(i)
		fmt.Println(nextInt(), nextInt())
	}

}
