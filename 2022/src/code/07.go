package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/8/29 14:05
 * @Desc:
 */

// 泛型

// 假设 T 是类型形参，在定义函数时它的类型是不确定的，类似占位符
/**

func Add(a T,b T) T {
	return a + b
}


 在上段代码中 T 被称为 类型参数 它不是具体的类型, 在定义函数时类型并不确定.
*/

// 这里类型约束使用了空接口，代表的意思是所有类型都可以用来实例化泛型类型 Queue[T] (关于接口在后半部分会详细介绍）
//
//type Queue[T interface{}] struct {
//	elements []T
//}
//
//// 将数据放入队列尾部
//
//func (q *Queue[T]) Put(value T) {
//	q.elements = append(q.elements, value)
//}
//
//// 从队列头部取出并从头部删除对应数据
//
//func (q *Queue[T]) Pop() (T, bool) {
//	var value T
//	if len(q.elements) == 0 {
//		return value, true
//	}
//
//	value = q.elements[0]
//	q.elements = q.elements[1:]
//	return value, len(q.elements) == 0
//}
//
//// 队列大小
//
//func (q Queue[T]) Size() int {
//	return len(q.elements)
//}

// Add 泛型函数

func Add[T int | float32 | float64](a T, b T) T {
	return a + b
}

// type any = interface{}
// [T any] 为类型约束 any 表示任意类型

func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v", v)
	}
	fmt.Println()
}

// 泛型 slice
type vector[T any] []T

// 泛型 map

type M map[string]any
type H[K string, V any] map[K]V

// 泛型 channel

type C chan any
type N[T any] chan any


// 泛型 类型约束

type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
	// ~ 代表底层数据类型
}
type Str interface {
	string
}
type NumStr interface {
	Num | Str
}


func add[T NumStr](a,b T) T {
	return a + b
}


func main() {
	//var i interface{} = 123
	////i.(int) // 类型断言
	//switch i.(type) {
	//case int:
	//	// do string conversion
	//}
	//fmt.Println(Add[int](1, 2))
	//fmt.Println(Add(1.0, 2.0))
	//printSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	//printSlice([]float64{1.1, 2.2, 3.3, 4.4, 5.5}) //1.1 2.2 3.3 4.4 5.5
	//printSlice([]string{"hello ", "world"})        //hello world
	//
	//vI := vector[int]{1, 2, 3, 4}
	//vS := vector[string]{"hello ", "world"}
	//fmt.Println(vI)
	//fmt.Println(vS)
	//
	//m := make(M)
	//m["String"] = "hello"
	//m["Int"] = 123
	//m["float"] = []float64{
	//	1.1,
	//	2.2,
	//	3.3,
	//}
	//fmt.Println(m)
	//// 另一种写法
	//h :=H[string,string]{
	//	"hello":"world",
	//}
	//h2 :=H[string,int] {
	//	"test":123,
	//}
	//fmt.Println(h)
	//fmt.Println(h2)
	//
	//ch :=make(C,2)
	//ch <- "hello"
	//ch <- 123456
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//
	//n2 := make(N[string],1)
	//n2 <- "world"
	//fmt.Println(<-n2)
	fmt.Println(add(10,20))
	fmt.Printf("%.1f\n", add(1.1, 2.2))
	fmt.Println(add("hello", " world")) //hello world

}
