package main

import "fmt"

// panic 又2个合理的用例
// 1. 发生一个不能恢复的错误, 此时程序不能继续运行. 一个例子就是 web 服务器无法绑定索要求的端口 在这种情况下, 就应该使用panic, 因为如果不能绑定端口 啥都做不了
// 2. 发生一个编程上的错误, 假如我们有一个接收指正参数的方法, 而其他人使用 nil 作为参数调用了它. 在这种情况下,使用 panic 因为这是一个编程错误,

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}

	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

// 发生panic 时的 defer
// 当函数发生 panic 时 会终止运行 在执行完所有的延迟函数后, 程序控制返回到该函数的调用方, 这样的过程会一直持续下去, 直到当协程的所有函数返回退出, 然后程序会打印出 panic 信息
// 接着打印出堆栈跟踪, 最后程序终止

func main() {
	defer fmt.Println("deferred call in fullName")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
