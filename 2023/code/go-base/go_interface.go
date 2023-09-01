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

type ZhiFuBao struct {
	// 支付宝
}

// Pay 支付宝的支付方法
func (z *ZhiFuBao) Pay(amount int64) {
	fmt.Printf("使用支付宝付款：%.2f元。\n", float64(amount/100))
}

// Checkout 结账
func Checkout(obj *ZhiFuBao) {
	//  支付100
	obj.Pay(100)
}

type WeChat struct {
}

// Pay 微信的支付方法
func (w *WeChat) Pay(amount int64) {
	fmt.Printf("使用微信付款：%.2f元。\n", float64(amount/100))
}

func main() {
	Checkout(&ZhiFuBao{})
}
