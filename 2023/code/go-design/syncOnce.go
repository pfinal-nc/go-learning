package main

import (
	"fmt"
	"sync"
)

// 单利模式

var once sync.Once

type singleOne struct {
}

var singleOneInstance *singleOne

func getInstanceOne() *singleOne {
	if singleOneInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now")
				singleOneInstance = &singleOne{}

			})
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleOneInstance
}
