/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-06-19 16:16:44
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// t := time.Now()
	// fmt.Print(t)
	// fmt.Println(t)
	// fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// t = time.Now().UTC()
	// fmt.Println(t) // Wed Dec 21 08:52:14 +0000 UTC 2011
	// fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// // calculating times:
	// week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	// week_from_now := t.Add(time.Duration(week))
	// fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// // formatting times:
	// fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	// fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
	// The time must be 2006-01-02 15:04:05

	p := fmt.Println
	// 得到当前时间
	now := time.Now()
	p(now)

	then := time.Date(
		2020, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 你可以提取出时间的各个组成部分
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 输出是星期一到日的 `Weekday` 也是支持的
	p(then.Weekday())

	// 这些方法来比较两个事件 分别测试一下是否是之前 之后或者是同一个时刻 精确到秒
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// 方法 `Sub` 可以计算两个时间之间的时间间隔
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 方法 `Add` 可以计算两个时间之间的时间间隔 使用 '-' 来将事件前移一个时间间隔
	p(then.Add(diff))
	p(then.Add(-diff))

}
