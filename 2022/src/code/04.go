package main

import "fmt"

// 浅拷贝 拷贝的是数据地址 只复制指向的对象的指针

func changeValue(m []int) {
	fmt.Printf("m=%v\taddr=%p \n", m, &m)
	m[2] = 8
}

func main() {
	arr := [4]int{1, 2, 3, 4}
	slice := arr[0:2]
	testSlice1 := slice
	testSlice2 := append(slice, 6)
	fmt.Printf("arr=%v\t addr=%p \n", arr, &arr)
	testSlice2 = append(testSlice2, 2, 3)
	testSlice2[2] = 5
	fmt.Printf("testSlice1=%v\t addr=%p \n", testSlice1, &testSlice1)
	fmt.Printf("testSlice2=%v\t addr=%p \n", testSlice2, &testSlice2)
	fmt.Printf("arr=%v\t addr=%p \n", arr, &arr)
	changeValue(testSlice2)  // 浅 拷贝  值传递 传递的
	fmt.Printf("testSlice2=%v\t addr=%p \n", testSlice2, &testSlice2) // 修改的是同一块内存地址
}
