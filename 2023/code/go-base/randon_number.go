package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

	// 数字解析
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 10, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
