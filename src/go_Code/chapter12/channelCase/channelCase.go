package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// Communication -------------------------------------------------------------------------------------------------------
// 协程间通讯
func Communication() {
	//可读可写通道
	ch := make(chan int, 0)
	go communicationFun1(ch)
	go communicationFun2(ch)
}

// 接收一个只写通道
func communicationFun1(ch chan<- int) {
	//	通过循环向通道内写入数据
	for i := 1; i <= 100; i++ {
		fmt.Println("只写通道内写入的数据：", i)
		ch <- i
		//为了更直观展现读写进度
		time.Sleep(time.Microsecond)
	}
}

// 接收一个只读通道
func communicationFun2(ch <-chan int) {
	for v := range ch {
		fmt.Println("只读通道内读出的数据：", v)
	}
}

// ConcurrentSync ------------------------------------------------------------------------------------------------------
// 并发场景下的同步机制
func ConcurrentSync() {
	//channel的缓冲容量要根据读写协程数量进行分配
	ch := make(chan int, 10)
	//开启两个协程并发的向channel内写入数据（无序）
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
	}()
	go func() {
		for i := 100; i <= 200; i++ {
			ch <- i
		}
	}()
	//	从channel中读取写入的数据并且打印
	go func() {
		sum := 0
		for value := range ch {
			fmt.Printf("读取出的第%v个数据，数据值为：%v\n", sum, value)
			sum++
		}
	}()
}

// NoticeAndMultiplexing -----------------------------------------------------------------------------------------------
// 通知协程退出与多路复用
func NoticeAndMultiplexing() {
	numch := make(chan int, 0)
	strCh := make(chan string, 0)
	done := make(chan struct{}, 0)
	go noticeAndMultiplexingFun1(numch)
	go noticeAndMultiplexingFun2(strCh)
	go noticeAndMultiplexingFun3(numch, strCh, done)
	time.Sleep(time.Second * 2)
	//关闭channel时候会向所有监听它的协程发送一个零值
	close(done)
}
func noticeAndMultiplexingFun1(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
}
func noticeAndMultiplexingFun2(ch chan<- string) {
	for i := 1; i <= 100; i++ {
		ch <- fmt.Sprintf("数字：%d", i)
	}
}
func noticeAndMultiplexingFun3(numch <-chan int, strch <-chan string, done <-chan struct{}) {
	i := 0
	for {
		//select子句作为一个整体阻塞，其中任意channel准备就绪就继续执行
		select {
		case num := <-numch:
			fmt.Println(num)
		case str := <-strch:
			fmt.Println(str)
		case <-done:
			fmt.Println("收到退出通知，退出当前协程")
			return
		}
		i++
		fmt.Println("累计执行次数：", i)
	}

}
func main() {
	//main函数是主协程
	//communicationFun1，communicationFun2是子协程
	//当主协程执行完毕后无论子协程是否执行完毕都退出
	//Communication()
	//ConcurrentSync()
	NoticeAndMultiplexing()
	//接收系统信号
	ch := make(chan os.Signal, 0)
	//接收收系统的中断操作或者杀死进程操作
	signal.Notify(ch, os.Interrupt, os.Kill)
	//如果接收到信号意味着进程退出
	<-ch
}
