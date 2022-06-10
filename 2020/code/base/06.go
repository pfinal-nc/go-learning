/*
 * ==================================
 * @Author: PFinal南丞
 * @Date: 2021-05-25 15:24:22
 * @Description:  高山仰止,景行行制,虽不能至,心向往之
 * ==================================
 */
package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	//"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
    fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
    fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
    fmt.Println("mem.NumGC:", mem.NumGC)
    fmt.Println("-----")
}

func main() {
	f, err := os.Create("/tmp/traceFile.out")
    if err != nil {
        panic(err)
    }
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
        fmt.Println(err)
        return
    }
    defer trace.Stop()
	var mem runtime.MemStats
    printStats(mem)
    for i := 0; i < 10; i++ {
        s := make([]byte, 50000000)
        if s == nil {
            fmt.Println("Operation failed!")
        }
    }
    printStats(mem)
    
}