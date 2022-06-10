/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-15 14:18:58
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 这里是一个基本的按照 RFC3339 格式解析时间的例子
	t := time.Now()

	// 时间解析使用同 Format 相同的形式值
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	p(t.Format(time.RFC3339))

	// `Format` 和 `Parse` 使用基于例子的形式来决定日期格式
	// 一般只要使用 time 包中提供的模式常量就行了, 但是也可以使用自定义的模式
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	// 对于纯数字表示的时间  可以使用标准的格式化字符串来提出事件值得组成
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d \n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// Parse 函数在输入的时间格式不正确时会返回一个错误
	ansic := "Mon Jan _2 15:04:05 2006"
	str, e := time.Parse(ansic, "8:41PM")
	p(e)
	p(str)

}
