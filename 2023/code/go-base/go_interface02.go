package main

//
//import "fmt"
//
//// 对于实现接口来说使用值接收者和使用指针接收者区别
//
//type Mover interface {
//	Move()
//}
//
//// 值接收者实现接口
//
//// 定义一个Dog 结构体 类型, 并使用接收者为 其定义一个 Move方法
//
//type Dog struct{}
//
//// Move 使用值接收者定义Move 方法实现 Mover 接口
//func (d Dog) Move() {
//	fmt.Println("狗会动")
//}
//
//func main() {
//	var x Mover
//	var d1 = Dog{} // d1 是Dog 类型
//	x = d1         // 可以将d1赋值给变量x
//	x.Move()
//
//	var d2 = &Dog{} // d2是Dog 指针类型
//	x = d2          // 也可以将d2赋值给变量 x
//	x.Move()
//}
//
//// 从上面的代码中可以发现, 使用值接收者实现接口之后, 不管是结构体类型 还是对应的结构体指针类型的变量都可以赋值给该接口变量
