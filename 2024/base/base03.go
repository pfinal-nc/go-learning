package main

import (
	"context"
	"fmt"
	"time"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/6/24
 * @Desc:
 * @Project: 2024
 */

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d: working\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 3; i++ {
		go worker(ctx, i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Main: cancelling context")
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Main: done")
}
