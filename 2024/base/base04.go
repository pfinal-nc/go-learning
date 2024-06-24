package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/6/24
 * @Desc:
 * @Project: 2024
 */

func main() {

	var s1 []int
	s2 := make([]int, 0)
	s4 := make([]int, 0)

	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)), *(*reflect.SliceHeader)(unsafe.Pointer(&s2)), *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}
