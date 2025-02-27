package main

import (
	"fmt"
	"time"
)

func PutNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func resNum(intChan chan int, resChan chan int, exitChan chan bool) {
	for {
		num_01, ok := <-intChan
		if !ok {
			break
		}
		flag := true
		for i := 2; i < num_01; i++ {
			if num_01%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resChan <- num_01
		}
	}
	fmt.Println("某一个协程因为读取不到数据退出")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 100)
	resChan := make(chan int, 8000) // 增加缓冲区大小
	exitChan := make(chan bool, 8)  //根据逻辑处理器个数

	startTime := time.Now()
	go PutNum(intChan)

	// 开启8个协程，如果是素数就放入到resChan中
	for i := 0; i < cap(exitChan); i++ {
		go resNum(intChan, resChan, exitChan)
	}

	go func() {
		for i := 0; i < cap(exitChan); i++ {
			<-exitChan
		}

		close(resChan)
	}()
	for res := range resChan {
		fmt.Println("素数有：", res)
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("程序执行时间: %v\n", elapsedTime)
	fmt.Println("主线程退出")
}
