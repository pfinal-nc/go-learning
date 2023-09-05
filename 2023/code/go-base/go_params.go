package main

import (
	"fmt"
	"time"
)

// panic 又2个合理的用例
// 1. 发生一个不能恢复的错误, 此时程序不能继续运行. 一个例子就是 web 服务器无法绑定索要求的端口 在这种情况下, 就应该使用panic, 因为如果不能绑定端口 啥都做不了
// 2. 发生一个编程上的错误, 假如我们有一个接收指正参数的方法, 而其他人使用 nil 作为参数调用了它. 在这种情况下,使用 panic 因为这是一个编程错误,

//func fullName(firstName *string, lastName *string) {
//	defer fmt.Println("deferred call in fullName")
//	if firstName == nil {
//		panic("runtime error: first name cannot be nil")
//	}
//
//	if lastName == nil {
//		panic("runtime error: last name cannot be nil")
//	}
//	fmt.Printf("%s %s\n", *firstName, *lastName)
//	fmt.Println("returned normally from fullName")
//}

// 发生panic 时的 defer
// 当函数发生 panic 时 会终止运行 在执行完所有的延迟函数后, 程序控制返回到该函数的调用方, 这样的过程会一直持续下去, 直到当协程的所有函数返回退出, 然后程序会打印出 panic 信息
// 接着打印出堆栈跟踪, 最后程序终止

//==============================

// recover
// recover 是一个内置函数 用于重新获的 panic 协程的控制
// recover 函数的标签如下所示:
// func recover() interface{}

//func recoverName() {
//	if r := recover(); r != nil {
//		fmt.Println("recovered from ", r)
//	}
//}
//
//func fullName(firstName *string, lastName *string) {
//	defer recoverName()
//	if firstName == nil {
//		panic("runtime error: first name cannot be nil")
//	}
//
//	if lastName == nil {
//		panic("runtime error: last name cannot be nil")
//	}
//	fmt.Printf("%s %s\n", *firstName, *lastName)
//	fmt.Println("returned normally from fullName")
//}
//
//func main() {
//	defer fmt.Println("deferred call in fullName")
//	firstName := "Elon"
//	fullName(&firstName, nil)
//	fmt.Println("returned normally from main")
//}

// 只有在相同的 go 协程中调用 recover 才管用, recover 不能恢复一个不同协程的 panic

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}

// 在上面的程序中 函数 b() 在发生 panic 函数 a()调用了一个延迟函数 recovery() 用于恢复 panic 函数 b() 作为一个不同的协程来调用, 下一行的Sleep 只是保证 a() 在b () 运行结束之后才退出
