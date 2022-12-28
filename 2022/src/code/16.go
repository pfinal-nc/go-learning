package main

import (
	"log"
	"os"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/12/16 17:54
 * @Desc:
 */

func main() {
	log.SetOutput(os.Stdout)
	log.Println("[4.426ms] [rows:1] SELECT * FROM `users` WHERE `id` = 1024")
	log.Printf("[GET] %d %s %s", 200, "OK", "/api/v1/users")
}
