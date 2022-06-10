/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-09 11:53:17
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
	"sort"
)

// 这里我们创建一个为内置的 []string 类型的别名为 ByLength 类型
type ByLength []string

// 在类型中 实现了  `sort.interface` 的 Len Less 和 swap 方法
// 这样 我们就可以使用 sort包的通用 sort 方法了
// Len 和 Swap 通常在各个类型中都差不多  Less 将控制实际的自定义排序逻辑

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {

	fruites := []string{"peach", "banner", "kiwi"}
	sort.Sort(ByLength(fruites))
	fmt.Println(fruites)
}
