/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-09 11:45:08
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	// 排序方法是针对内置数据类型的
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// 也可以使用 sort 来检查一个序列是不是已经是排好序的
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
