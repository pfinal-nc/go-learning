/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-16 11:43:02
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// 开始 这里是展示如何写入一个字符串到一个文件
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("dat1", d1, 0644)
	check(err)

	// 对于更细粒度的写入,你可以使用带缓冲的写入器

	f, err := os.Create("dat2")
	check(err)

	defer f.Close()

	// 写入字节切片
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString 也是可用的
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用Sync 来缓冲区的信息写入磁盘
	f.Sync()

	// bufio 提供了带缓冲的读取器和写入器
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	// 使用 `Flush` 来确保所有缓存的操作已写入底层写入器
	w.Flush()
}
