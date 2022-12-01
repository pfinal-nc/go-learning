package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/12/2 00:20
 * @Desc:
 */

type Person1 struct {
	age int
}

func main(){
	person :=&Person1{28}

	defer fmt.Println(person.age)

	defer func(p *Person1) {
		fmt.Println(p.age)
	}(person)

	defer func() {
		fmt.Println(person.age)
	}()
	person.age = 29
}