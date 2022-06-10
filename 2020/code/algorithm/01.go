// 题目:古典问题:有一对兔子,从出生后第 3 个月起每个月都生一对兔子,小兔 子长到第三个月后每个月又生一对兔子,假如兔子都不死,问每个月的兔子总数为多少?
package main

import (
	"fmt"
)

func main() {
	f1,f2 :=1,1
	month := 6
	for i :=0; i < month;i++ {
		f1, f2 = f2, f1+f2
		fmt.Printf("%d===F1--1\n", f1)
		fmt.Printf("%d===F2--2\n", f2)
	}
}