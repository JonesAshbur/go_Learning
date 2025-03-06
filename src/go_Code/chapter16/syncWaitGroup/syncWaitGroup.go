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
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func waitGroupCase1() {
	//单协程
	var a, b int = 1, 2
	start := time.Now()
	for i := 0; i < 10000000000; i++ {
		multi(a, b)
	}
	t := time.Since(start)
	fmt.Println("单协程执行时间：", t)
	start = time.Now()
	//多协程
	wg := sync.WaitGroup{}
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1250000000; j++ {
				multi(a, b)
			}
		}()
	}
	wg.Wait()
	t = time.Since(start)
	fmt.Println("多协程执行时间：", t)
}
func multi(a, b int) int {
	return a * b
}

func waitGroupCase2() {
	ch := make(chan []int, 1000)
	start := time.Now()
	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		i := 0
		for item := range ch {
			fmt.Println(multi(item[0], item[1]))
			i++
		}
		time.Sleep(time.Second * 3)
		fmt.Println("数据处理完成，数据条数：", i)
	}()
	wg := &sync.WaitGroup{}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		wg2.Add(1)
		go func(wg1 *sync.WaitGroup) {
			defer wg.Done()
			defer wg2.Done()
			for j := 1; j <= 500; j++ {
				ch <- []int{i, j}
			}
		}(wg)
	}
	wg.Wait()
	close(ch)
	wg2.Wait()
	t := time.Since(start)
	fmt.Println("花费时间：", t)
}
func main() {
	waitGroupCase1()
	waitGroupCase2()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
