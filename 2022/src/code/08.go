package main

import (
	"fmt"
	"strconv"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/8/30 10:38
 * @Desc:
 */

// 泛型 方法约束

type Price int

type ShowPrice interface {
	String() string
}

func (p Price) String() string {
	return strconv.Itoa(int(p))
}

func ShowPriceList[T ShowPrice](s []T) []string {
	var res []string
	for _, v := range s {
		res = append(res, v.String())
	}
	return res
}

func main() {
	fmt.Println(ShowPriceList([]Price{1, 2, 3, 4, 5}))
}
