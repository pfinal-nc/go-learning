package main

import (
	"fmt"
)
import "github.com/gomodule/redigo/redis"

var conn redis.Conn

func init() {
	var err error
	conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}

}
func ListPush(args ...interface{}) (num int64) {
	v, err := conn.Do("RPUSH", args...)
	if err != nil {
		fmt.Println(err)
	}
	num = v.(int64)
	return
}

func ListPop(key string) (v string) {
	v, err := redis.String(conn.Do("LPOP", key))
	if err != nil {
		fmt.Println(err)
	}
	return
}

func main() {
	fmt.Println(ListPush("list5", "a", "b", "c"))
	fmt.Println(ListPop("list5"))
}
