// 判断 101-200 之间有多少个素数,并输出所有素数。
package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0 
	for i :=101; i <= 200; i++ {
		mid := int(math.Sqrt(float64(i)))
		fmt.Println(mid)
		for j :=2; j <= mid; j++ {
			if i%j == 0 {
				count = count + 1
				break
			}
		}
	}
	fmt.Println(100 - count)
}