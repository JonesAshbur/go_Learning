package main

import (
	"fmt"
)

func WriteData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("写入数据：", i)
	}
	close(intChan)
}
func ReadData(intChan chan int, exitChan chan bool) {
	for {
		value, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读出数据：", value)
	}
	exitChan <- true
	close(exitChan)
}
func main() {

	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go WriteData(intChan)
	go ReadData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
	fmt.Println("程序结束")

	// 如果一直向channel写入数据却不读取，则会发生阻塞（读取慢，读写频率不同都不会阻塞）
}
