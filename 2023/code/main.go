package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	if v, ok := m["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("aaaaa")
	}
}
