package main

import (
	"fmt"
	"os"
)

// fmt 的几种输出

// Print 系列函数将内容输出到系统的标准输出, 区别在于 print 函数直接输出内容. Printf 函数支持格式化输出字符串 Println 函数会在输出内容的结尾添加一个换行符

// Fprint 系列函数会将内容输出到一个 io.Writer 接口类型的变量 w 中, 通常这个函数往文件中写入内容
// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

func main() {
	// fmt.P 系列
	//fmt.Print("PFinalClub")
	//name := "PFinalClub"
	//fmt.Printf("我是: %s\n", name)
	//fmt.Println("##########")

	fmt.Fprintf(os.Stdout, "向标准版输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}
	name := "PFinalClub"
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}
