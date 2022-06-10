/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-11-04 14:17:28
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (stu *Student) Hello(person string) string {
	return fmt.Sprintf("Hello %s, I am %s", person, stu.name)
}

func main() {
	stu := Student{
		name: "PFinal",
		age:  18,
	}
	fmt.Println(stu.Hello("World"))

	stu2 := new(Student)
	fmt.Println(stu2.Hello("World"))

}
