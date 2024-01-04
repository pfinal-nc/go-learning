package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/1/3
 * @Desc:
 * @Project: 2023
 */

func main() {
	fmt.Println(rand.Intn(100), ",")
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Println((rand.Float64() * 5) + 5)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100), ",")
	fmt.Println(r1.Intn(100))
}
