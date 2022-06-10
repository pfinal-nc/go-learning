/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-15 16:27:34
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("dat")
	check(err)
	fmt.Print(string(dat))

	fmt.Println("\n\n")
	f, err := os.Open("dat")
	check(err)

	// 从文件开始位置读取一些字节  这里最多读取5个字节 并且 这也是我们实际读取的字节数
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 可以 Seek 到一个 文件中已知的位置从这个位置开始进行读取
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// io 包提供了一些可以帮助我们进行文件读取的函数
	// 例如  上面的读取可以使用 ReadAtLeast 函数 得到一个更健壮的实现
	o3, err := f.Seek(6, 0)
	fmt.Printf("%d\n", o3)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 没有内置的回转支持  但是 使用 Seek(0, 0) 实现
	_, err = f.Seek(0, 0)
	check(err)

	// bufio 包实现了带缓冲的读取  这不仅对于那些小文件比较有用
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(100)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 任务结束后要关闭文件 操作后立即使用 defer 关闭文件
	f.Close()
}
