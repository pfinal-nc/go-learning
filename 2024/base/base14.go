package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/8/22
 * @Desc:
 * @Project: 2024
 */

func SSEHandler(w http.ResponseWriter, r *http.Request) {
	// 设置SSE相关的响应头
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 检查是否支持Flush
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// 创建一个通道，用于将输入的数据发送到SSE
	inputChan := make(chan string)

	// 启动一个Goroutine来读取标准输入并发送到通道
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Println("Read from stdin:", text) // 输出读取到的内容
			inputChan <- text
		}
		close(inputChan)
	}()

	// 监听通道中的数据并推送到客户端
	for {
		select {
		case msg, ok := <-inputChan:
			if !ok {
				// 通道关闭，结束SSE
				fmt.Fprint(w, "data: Connection closed\n\n")
				flusher.Flush()
				return
			}
			fmt.Println("Pushing to client:", msg) // 输出即将推送的内容
			_, err := fmt.Fprintf(w, "data: %s\n\n", msg)
			if err != nil {
				// 推送失败，可能是客户端断开了连接
				fmt.Println("Client disconnected:", err)
				return
			}
			flusher.Flush()
		}
	}
}

func main() {
	http.HandleFunc("/events", SSEHandler)

	fmt.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
