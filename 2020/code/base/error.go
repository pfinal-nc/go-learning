/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-06-08 18:03:18
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"errors"
	"fmt"
)

//var errNotFound error = errors.New("Not found error")
// 错误通常是最后一个返回值并且是 `error` 一个内建的接口

func f1(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` 构造一个使用给定的错误信息的基本 `error` 值
		return -1, errors.New("can't work with 42")

	}
	return arg + 3, nil
}

// 可以通过实现 `Error` 方法来自定义 `error`类型

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {

	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	//fmt.Printf("%+v\n", errNotFound)
	for _, i := range []int{7, 42} {
		r, e := f1(i)
		if e == nil {
			fmt.Println("f1 worked:", r)
			continue
		}
		fmt.Println("f1 failed:", e)
	}

	for _, j := range []int{7, 42} {
		if r, e := f2(j); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
}
