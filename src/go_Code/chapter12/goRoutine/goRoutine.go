package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func test() {
	for i := 0; i <= 10; i++ {
		fmt.Println("test函数hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {

	// 线程和进程
	// 进程就是程序在操作系统中的一次执行过程，是系统资源分配和调度的基本单位
	// 线程是进程的一个执行实例，是程序执行的最小单元，是比进程更小的能独立运行的基本单位
	// 一个进程可以创建销毁多个线程，一个进程中的多个线程可以并发执行
	// 一个程序至少有一个进程，一个进程至少有一个线程

	// 多线程程序在单核上运行---->并发
	// 多线程程序在多核上运行---->并行

	// go主线程，有人直接成为线程，也可以理解为进程，一个go线程上可以起多个协程，协程是轻量级的线程

	// go协程的特点：
	// 1.有独立的栈空间
	// 2.共享程序堆空间
	// 3.调度由用户控制
	// 4.协程是轻量级的线程

	// 注意事项：
	// 1.如果主线程执行完毕退出了，协程即使还没有执行完毕也会退出
	// 2.协程也可以执行完毕先行退出
	// 3.主线程是物理态，直接作用于CPU，协程是从主线程开启的，是逻辑态轻量级线程，对资源消耗较小
	// 4.协程机制是golang的重要特点，其他编程语言的并发机制是基于线程的，开启过多线程耗费资源较大
	go test() //开启协程
	for i := 0; i <= 10; i++ {
		fmt.Println("main函数中hello golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

	// MPG调度模型
	// M：操作系统主线程
	// P：协程执行需要的上下文
	// G：协程

	num := runtime.NumCPU()
	fmt.Println("num=", num)
	// runtime.GOMAXPROCS(num - 1)设置cpu数量运行go程序
	// go1.8以后默认让cpu运行在多核
}
