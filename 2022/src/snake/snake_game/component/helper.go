package component

import (
	"fmt"
	"time"
)

// Catch 全局捕获异常
func Catch() {
	if err := recover(); err != nil {
		//打印异常
		fmt.Println(fmt.Sprintf("Catch faild:%s", err))
	}
}

// Run 执行方法
func Run(f func()) {
	defer Catch()
	f()
}

//flush 刷新
func flush(score int) {
	time.Sleep(time.Duration(150-(score/10)) * time.Millisecond)
}
