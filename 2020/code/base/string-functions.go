/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-10 14:19:21
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:  ", s.Count("test", "t"))
	p("HasPrefix:  ", s.HasPrefix("test", "t"))
	p("HasSuffix:  ", s.HasSuffix("test", "s"))
	p("Index:  ", s.Index("test", "e"))
	p("Join:  ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:  ", s.Repeat("*", 5))
	p("Replace:  ", s.Replace("fooo", "o", "0", -1))
	p("Split:	", s.Split("a-b-c-d-e", "-"))
	p("Join: ", s.Join(s.Split("a-b-c-d-e", "-"), "*"))
	p("ToLower: ", s.ToLower("TEST"))
	p("ToUpper: ", s.ToUpper("test"))

	p("Len: ", len("hello"))
	p("Char: ", "hello"[1])

}
