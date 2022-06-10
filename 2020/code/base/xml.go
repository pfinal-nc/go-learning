/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-15 10:10:59
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	//fmt.Println(coffee)

	// 使用 MarshalIndent 函数将结构体转换为 XML 格式
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	// fmt.Println(string(out))

	// 明确的为输出结果添加一个通用的 XML 头部信息
	fmt.Println(xml.Header + string(out))

	// 使用 Unmarshal 函数将 XML 格式转换为结构体
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	// parent>child>plant  字段标签告诉编码器嵌套
	type Nesting struct {
		XMLName xml.Name `xml:"parent"`
		Plants  []*Plant `xml:"child>plant"`
	}

	netsting := &Nesting{}
	netsting.Plants = []*Plant{coffee, tomato}
	out, _ = xml.MarshalIndent(netsting, " ", "  ")
	fmt.Println(string(out))
}
