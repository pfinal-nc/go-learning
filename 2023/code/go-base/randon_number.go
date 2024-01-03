package main

import (
	"fmt"
	"math/rand"
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
}
