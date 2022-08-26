package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/8/26 16:17
 * @Desc:
 */

// 值传递 如果函数传参是值类型, 会复制值的副本. 如果函数传参是引用类型, 就会复制这个引用指针的值

type Person struct {
	Name string
	Age  int64
}

func modifiedAge(per Person) {
	fmt.Printf("函数里接收到 struct 的内存地址是: %p\n", &per)
	per.Age = 10
}

func main() {
	per := Person{
		Name: "liu",
		Age:  int64(8),
	}
	fmt.Printf("原始struct地址是: %p\n", &per)
	modifiedAge(per)
	fmt.Println(per)
}
