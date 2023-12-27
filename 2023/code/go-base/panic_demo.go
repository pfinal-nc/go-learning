package main

import "os"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/27
 * @Desc:
 * @Project: 2023
 */
// panic 意味着有些出乎意料的错误发生, 通常用它来表示程序正常运行中不应该出现的

func main() {
	panic("a problem")
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
