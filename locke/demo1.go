package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sy := &sync.Mutex{}
	go test(sy)
	go test(sy)
	go test(sy)
	time.Sleep(4 * time.Second)
}
func test(sy *sync.Mutex) {
	sy.Lock()
	fmt.Println("开始刮痧！")
	fmt.Println("刮痧结束！")
	time.Sleep(time.Second)
	sy.Unlock()
}
