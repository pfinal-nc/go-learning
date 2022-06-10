/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-09-10 15:49:50
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */

package main

import (
	"encoding/json"
	"fmt"
)

// 定义2个结构体 来编码和解码
type Response struct {
	Page   int
	Fruits []string
}

type ResponseA struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	strB, _ := json.Marshal(2.34)
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	res1D := &Response{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 你可以给结构字段声明标签来自定义编码的 JSON 数据键名称
	res2D := &ResponseA{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b","c"]}`)
	fmt.Println(byt)

	var dat map[string]interface{}

	// 这里就是实际的解码和相关的错误检查
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 为了使用解码 map 中的值  我们需要将他们进行适当的类型转换
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	fmt.Println(strs)

}
