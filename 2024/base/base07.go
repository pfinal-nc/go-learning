package main

import (
	"os"
	"os/exec"
	"syscall"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/6/26
 * @Desc:
 * @Project: 2024
 */
// 这里 exec 同样需要使用环境变量, 所以我们使用 os.Environ() 来获取环境变量
// 如果这个调用成功,那么我们的进程将在这里被替换成成/bin/ls -a -l -h 进程.如果存在错误,那么我们将抛出一个 panic
func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
