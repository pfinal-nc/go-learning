package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/11/14 12:32
 * @Desc:
 */

func main() {
	const (
		a1  = iota
		a2
		a3
	)
	a :=100
	p := &a
	fmt.Println(a1,a2,a3)
	fmt.Printf("%T\n",a1)
	fmt.Printf("%T\n",p)

	c :=[]int{1,2,3}
	fmt.Printf("%T\n",c)
}