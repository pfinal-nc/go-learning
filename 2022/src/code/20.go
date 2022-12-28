package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/12/28 23:01
 * @Desc:
 */

type Students struct {
	id    uint
	name  string
	male  bool
	score float64
}

func NewStudent(id uint, name string, score float64) *Students {
	return &Students{id: id, name: name, score: score}
}

func (s *Students) SetName(name string) *Students {
	s.name = name
	return  s
}

func (s Students) String() string {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}",
		s.id, s.name, s.male, s.score)
}

func main() {
	student := NewStudent(1, "PFinal南丞", 100)
	fmt.Println(student)
	student.SetName("彭儿")
	fmt.Println(student)
}
