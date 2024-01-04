package main

import (
	"fmt"
	"net/url"
	"strings"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/1/4
 * @Desc:
 * @Project: 2023
 */

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	// l := "https://friday-go.icu"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	p, _ := u.User.Password()
	fmt.Println(p)
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h)
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
