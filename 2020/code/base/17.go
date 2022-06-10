/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-11-03 12:04:59
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Uid int
	Tid int
}

func main() {
	var arr []User
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var uid int = 3
	for _, v := range a {
		user := User{Uid: uid, Tid: v}
		arr = append(arr, user)
	}
	b, err := json.Marshal(arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
