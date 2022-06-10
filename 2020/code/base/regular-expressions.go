/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-10 15:33:33
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"regexp"
)

func main() {

	// 检测字符串是否匹配正则表达式
	match, _ := regexp.MatchString("p([a-z]+)", "pfinal")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)l")
	// 检测字符串是否匹配正则表达式
	fmt.Println(r.MatchString("pfinal"))

	// 查找字符串中是否匹配正则表达式

	fmt.Println(r.FindString("pfinal per"))

	//查找第一次匹配的字符串返回匹配开始的位置和结束的位置
	fmt.Println(r.FindStringIndex("pfinal pterl"))

	// 并且返回所有的子匹配
	fmt.Println(r.FindStringSubmatch("pfinal pterl"))

	// 类似 返回完全匹配和局部匹配的索引位置
	fmt.Println(r.FindStringSubmatchIndex("pfinal pterl"))

	fmt.Println(r.FindAllString("pfinal pterl pincl", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("pfinal pterl pincl", -1))
}
