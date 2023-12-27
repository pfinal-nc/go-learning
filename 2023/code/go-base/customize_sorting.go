package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/27
 * @Desc:
 * @Project: 2023
 */

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"apple", "banana", "pear", "cherry", "date"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
