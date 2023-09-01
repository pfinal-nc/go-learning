package main

import "fmt"

//
//type Sayer interface {
//	Say()
//}
//
//func MakeHungry(s Sayer) {
//	s.Say()
//}
//
//type Cat struct{}
//
//func (c Cat) Say() {
//	fmt.Println("喵喵喵")
//}
//
//func main() {
//	var c Cat
//	MakeHungry(c)
//}

// 接口类型变量
// var x Sayer // 声明一个 Sayer 类型的变量 x
// a := Cat{}  声明一个Cat类型变量a
// b := Dog{}  声明一个Dog类型变量b
// x = a 可以把Cat类型变量直接赋值给x
// x.Say()
// x=b 可以把Dog 类型变量直接赋值给x
// x.Say()

// Mover 值接收者和指针接收者
// Mover 定义一个接口类型
//type Mover interface {
//	Move()
//}
//
//// 值接收者实现接口
//
//// Cat 猫结构体类型
//type Cat struct{}
//
//// Move 使用指针接收定义Move方法实现Mover接口
//func (c *Cat) Move() {
//	fmt.Println("猫会动")
//}
//
//// 此时实现 Mover 接口是 *Cat 类型, 可以将 *Cat 类型的变量直接复制给 Mover 接口类型的变量x
//
//var c1 = &Cat{} // c1是*Cat类型
//x = c1          // 可以将c1当成Mover类型
//x.Move()

// 类型与接口的关系
// 一个类型可以同时实现多个接口, 而接口间彼此独立, 不知道对方的实现。

// Sayer 接口
type Sayer interface {
	Say()
}

// Mover 接口
type Mover interface {
	Move()
}

// Dog 既可以实现 Sayer 接口 也可以实现 Mover 接口
type Dog struct {
	Name string
}

// Say 实现 Sayer 接口
func (d Dog) Say() {
	fmt.Printf("%s会叫汪汪汪\n", d.Name)
}

// Move 实现 Mover 接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

func main() {
	// 同一个类型实现不同的接口互相不影响使用
	var d = Dog{Name: "PFinal"}

	var s Sayer = d
	var m Mover = d
	s.Say()
	m.Move()
}
