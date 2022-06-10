/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-16 19:14:43
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// 使用 Join 构造路径
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println(p)
	// 使用 Join 代替手动拼接路径
	fmt.Println(filepath.Join("dir//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir2", "filename"))
	// dir 和 base 可以呗用于分割路径中的目录和文件
	fmt.Println("Dir(p)", filepath.Dir(p))
	fmt.Println("base(p)", filepath.Base(p))

	// 判断路径是否为绝对路径
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"
	// 查看扩展名
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// 想获取文件名，可以使用 filepath.Base
	fmt.Println(strings.TrimSuffix(filename, ext))

	// 如果相对路径不存在 则返回错误
	// rel 寻找 basepath 与 targpath 之间的相对路径
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
