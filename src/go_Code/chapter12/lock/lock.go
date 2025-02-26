package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	resMap = make(map[int]int, 10)
	lock   sync.Mutex
)

func calFun(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
	}
	lock.Lock()
	resMap[n] = res
	lock.Unlock()
}
func main() {

	// 开启多个协程
	// 在运行某个程序时，使用go build -race可以知道是否存在资源竞争
	// 通过全局变量互斥锁解决资源竞争
	for i := 1; i < 20; i++ {
		go calFun(i)
	}
	time.Sleep(time.Second * 3)
	lock.Lock()
	for index, value := range resMap {
		fmt.Printf("resMap[%v]=%v\n", index, value)
	}
	lock.Unlock()
}
