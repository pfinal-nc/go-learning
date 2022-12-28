package main

import (
	"fmt"
	"os"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/12/27 16:47
 * @Desc:
 */

func main() {
	err := os.Mkdir("text",0755)
	if err != nil {
		panic(err)
	}else {
		fmt.Println("/tmp/test_go_main_dir has been created")
	}

}