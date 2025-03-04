// sync.WaitGroup作用
//1.等待一组协程完成
//2.工作原理：通过计数器来获取协程的完成情况
//3.启动一个协程时计数器+1，协程退出时计数器-1
//4.通过wait方法阻塞主协程，等待计数器清零后才能继续执行后续操作

// sync.WaitGroup应用场景
// 通过协程并行执行一组任务，且任务全部完成后才能进行下一步操作的情况

// sync.WaitGroup注意事项
// 协程间传递时需要以指针的方式或者是闭包的方式引用WaitGroup对象，否则死锁

// sync.Con作用
// 1.设置一组协程根据条件阻塞，可以根据不同的条件阻塞
// 2。可以根据条件唤醒相应的协程

// sync.Con应用场景
// 应用于一发多收的场景，即一组协程需要等待某一个协程完成一些前置准备情况

// sync.Con注意事项
// 被叫方必须有锁
// 主叫方可以持有锁也可以不持有锁
// 尽可能减少无效唤醒

// 协程并发注意事项：不能有锁，日志打印，I/O操作（串行）
package main

import (
	"fmt"
	"sync"
	"time"
)

func waitGroupCase() {
	var a, b int = 1, 2
	start := time.Now()
	for i := 0; i < 10000000000; i++ {
		multi(a, b)
	}
	t := time.Since(start)
	fmt.Println(t)
	start = time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 8; i++ {
		wg.Add(1)
	}
}
func multi(a, b int) int {
	return a * b
}
func main() {
	waitGroupCase()
}
