package main

import (
	"fmt"
	"github.com/fatih/structs"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/2/28 17:24
 * @Desc:
 */

// UserInfo 用户信息
type UserInfo struct {
	Name    string `json:"name" structs:"name"`
	Age     int    `json:"age" structs:"age"`
	Profile `json:"profile" structs:"profile"`
}

type Profile struct {
	Hobby string `json:"hobby" structs:"hobby"`
}

func main() {

	u1 := UserInfo{Name: "测试", Age: 10, Profile: Profile{"双色球"}}
	m3 := structs.Map(&u1)
	for k, v := range m3 {
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
}
