package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/7/23 11:27
 * @Desc:
 */
/**
chan是什么类型就能存放什么类型的数据，我们看看反射，他可不管你什么类型，随便传吧,fmt.Println中传入任何类型的数据都可以打印
1) 反射可以在运行时动态获取变量的各种信息, 比如变量的类型(type),类别(kind)
reflectType(变量名)　reflectValue(变量名)
2) 如果是结构体变量,还可以获取到结构体本身的信息(包括结构体的字段、方法)
3) 通过反射,可以修改变量的值,可以调用关联的方法。
 */

func reflectTest01(b interface{})  {
	// 通过反射获取传递进来的变量的类型
	// 1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp: ", rTyp)
	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(180)
	fmt.Println("rVal: ", rVal)
}

func main() {
	var num int = 100
	reflectTest01(&num)
	fmt.Println("num: ", num)
}
