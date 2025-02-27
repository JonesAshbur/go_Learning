package main

import (
	"fmt"
	"time"
)

func Print() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello golang")
	}
}
func test() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("test函数发生异常:", error)
		}
	}()
	var demoMap map[int]string
	demoMap[0] = "hello golang" // 没有make
}
func main() {
	// 演示recover捕获panic
	go Print()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("主函数执行顺序：", i)
		time.Sleep(time.Second)
	}
}
