/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-15 15:59:19
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 例如 `rand.Intn` 返回一个随机的整数 n
	// `0 <= n <=1.0`
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
