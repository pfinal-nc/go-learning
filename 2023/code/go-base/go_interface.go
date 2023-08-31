package main

import "fmt"

type Sayer interface {
	Say()
}

func MakeHungry(s Sayer) {
	s.Say()
}

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵")
}

func main() {
	var c Cat
	MakeHungry(c)
}
