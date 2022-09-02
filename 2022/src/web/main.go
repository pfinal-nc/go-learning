package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
	"net/http"
	"strings"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/9/1 10:38
 * @Desc:
 */

type MyHandler struct {
}

func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloGolang(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloGolang(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello Golang!")
}

func dockerActionContainerList(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		_, _ = fmt.Fprintf(w, container.ID + "\n")
	}

}

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	} // 解析参数
	fmt.Println(r.Form)
	fmt.Println("URL:", r.URL.Path)
	fmt.Println("Scheme:", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println(k, ":", strings.Join(v, ":"))
	}
	_, _ = fmt.Fprintf(w, "你好,PFinalClub")
}

func main() {
	http.HandleFunc("/", sayHelloWorld)
	http.HandleFunc("/docker", dockerActionContainerList)
	//handler := MyHandler{}
	err := http.ListenAndServe("127.0.0.1:9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
