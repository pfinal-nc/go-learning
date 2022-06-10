/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-11-04 14:37:15
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

type Person interface {
	getName() string
}

type Stu struct {
	name string
	age  int
}

func (stu *Stu) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func main() {
	var p Person = &Stu{
		name: "Tom",
	}
}
