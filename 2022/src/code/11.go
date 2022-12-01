package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/12/2 00:44
 * @Desc:
 */

func main() {
	a := [...]int{0,1,2,3,4,5,6,7,8,9}

	s :=a[2:8:8]
	fmt.Println(s)
	fmt.Println(len(s),cap(s))
}
