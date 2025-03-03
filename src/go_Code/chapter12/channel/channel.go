package main

import (
	"fmt"
)

func main() {

	// 为什么需要channel
	// 通过设置延时执行时间很难断定goroutine全部完成时间
	// 主线程休眠时间长会增加等待时间，等待时间太短会打断协程的执行
	// 通过全局变量加互斥锁实现通讯，不利于多个协程对全局变量的读写操作

	// channel的本质是队列
	//协程间通信，即协程间数据传递
	//并发场景下利用channel的阻塞机制作为同步机制
	//利用channel关闭时发送广播的特性，作为协程退出通知
	// 线程安全，多协程访问时，无需加锁，不会发生资源竞争问题
	// channel有类型，例如int类型的channel只能存放int，interface{}空接口channel可以接受多个类型
	// channel是引用类型，必须初始化才能写入数据（先make）
	// channel的容量不能自动增长
	// channel写入数据量不能超过容量
	// 在没有使用协程的情况下，如果数据已经全部取出，再取就会报告deadlock
	var intChan chan int = make(chan int, 2)
	fmt.Println(intChan)
	intChan <- 10
	var num_01 int = 10
	intChan <- num_01
	fmt.Printf("intChan长度：%v,intChan容量：%v\n", len(intChan), cap(intChan))
	num_02 := <-intChan
	fmt.Println("取出的数据：", num_02)
	<-intChan //可以不定义接收变量
	fmt.Printf("intChan长度：%v,intChan容量：%v\n", len(intChan), cap(intChan))

	// channel的关闭
	// 使用内置函数close可以关闭channel，关闭后不可以向channel中写入数据，但是可以读取

	demoChan := make(chan int, 3)
	demoChan <- 1
	demoChan <- 2
	close(demoChan)
	num_03 := <-demoChan
	fmt.Println("channel关闭后读取的数据：", num_03)
	// channel的遍历
	// for range遍历
	// 1.遍历channel时如果没有关闭channel，会报错deadlock
	// 2.如果已经关闭channel，会正常遍历数据，遍历完成退出遍历
	// 3.遍历channel不能使用普通的for循环，不能使用len(channel)，cap(channel)作为循环次数
	intChan_01 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan_01 <- i * 2
	}
	close(intChan_01)
	for value := range intChan_01 {
		fmt.Println("value:", value)
	}

	// channel注意事项：
	// 1.channel可以声明成只读或者只写的形式
	// 默认情况下管道是双向的
	// 只写：var chan_01 chan<- int
	// 只读：var chan_02 <-chan int
	// 2.使用select可以解决从管道取数据的阻塞问题
	// 3.go协程中使用recover，解决写成中出现panic，导致程序崩溃问题（即使协程出现问题，主线程不受影响）

	intChan_02 := make(chan int, 10)
	for i := 0; i < cap(intChan_02); i++ {
		intChan_02 <- i
	}
	charChan_01 := make(chan string, 5)
	for i := 0; i < cap(charChan_01); i++ {
		charChan_01 <- "hello" + fmt.Sprintf("%d", i)
	}
	for {
		select {
		case value := <-intChan_02: //如果channel没有关闭，不会一直阻塞导致死锁，会继续执行后面的case
			fmt.Println("从intChan_02读取了数据：", value)
		case value := <-charChan_01:
			fmt.Println("从charChan_01读取了数据：", value)
		default:
			// time.Sleep(time.Second * 2)
			fmt.Println("都没取到")
			return
		}
	}
}
