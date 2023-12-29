package main

import (
	"fmt"
	"regexp"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/29
 * @Desc:
 * @Project: 2023
 */

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ub", "pfinalclub")
	fmt.Println(match)
	r, _ := regexp.Compile("p([a-z]+)l")
	fmt.Println(r.MatchString("pfinalclub"))
	fmt.Println(r.FindString("pfinalclub"))
	fmt.Println(r.FindStringIndex("pfinal"))
	fmt.Println(r.FindStringSubmatch("pfinalclub")) // Submatch 返回完全匹配和局部匹配的字符串
	fmt.Println(r.FindStringSubmatchIndex("pfinalclub"))
	fmt.Println(r.FindAllString("pfinalclub", -1))

}
