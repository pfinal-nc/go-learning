/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-10 10:33:21
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
	"strings"
)

// 返回目标字符串 `t` 出现的第一个所以位置
// 或者在没有匹配值时返回 -1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 如果目标字符串 `t` 在字符串切片 `vs` 中出现,则返回 `true`
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 如果这些切片中的字符串都不包含 `t` 则返回 `true`
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 如果切片的所有字符串都包含 `t` 则返回 `true`
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// 返回一个包含所有切片中满足条件 f 字符串的新切片

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 返回一个对原始切片中所有字符串执行函数 f 后的新切片
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "pears"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(s string) bool {
		return strings.Contains(s, "a")
	}))

	// 上面的例子都是用的匿名函数，但是你也可以使用类型正确的命名函数
	fmt.Println(Map(strs, strings.ToUpper))
}
