/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-06-15 10:28:26
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import "fmt"

// 递归函数
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
	// anArray := [5]int{0, 1, -1, 2, -2}
	// for i, value := range anArray {
	// 	fmt.Println("index:", i, "value:", value)
	// }

}
