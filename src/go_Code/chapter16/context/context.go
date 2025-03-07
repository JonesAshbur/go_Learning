// context(上下文)
// 用于进程之间信息和信号传递
// 用于服务之间信息和信号传递
// 应用场景：
// 1.父子协程间取消信号传递
// 2.用于客户端和服务器之间的信号传递
// 3.用于设置请求超时时间
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

// 传统写法
func traditionalCase() {
	done := make(chan struct{})
	go fun1(done)
	go fun1(done)
	time.Sleep(time.Second * 2)
	close(done)
}
func fun1(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("协程退出")
			return
		}
	}
}

// 使用context
func contextCase() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "desc", "contextCase")
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	data := [][]int{
		{1, 2},
		{4, 5},
	}
	ch := make(chan []int)
	go calculate(ctx, ch)
	for i := 0; i < len(data); i++ {
		ch <- data[i]
	}
	time.Sleep(time.Second * 10)
}
func calculate(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			ctx := context.WithValue(ctx, "desc", "calculate")
			ch1 := make(chan []int)
			go sumContext(ctx, ch1)
			ch1 <- item
			ch2 := make(chan []int)
			go multiContext(ctx, ch2)
			ch2 <- item
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("calculate 协程退出,context desc:%s,错误消息：%s\n", desc, ctx.Err())
			return
		}
	}
}
func sumContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := sum(a, b)
			fmt.Printf("%d + %d = %d\n", a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("sumContext 协程退出,context desc:%s,错误消息：%s\n", desc, ctx.Err())
			return
		}
	}
}
func multiContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := multi(a, b)
			fmt.Printf("%d * %d = %d\n", a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("multiContext 协程退出,context desc:%s,错误消息：%s\n", desc, ctx.Err())
			return
		}
	}
}
func sum(a, b int) int {
	return a + b
}
func multi(a, b int) int {
	//time.Sleep(time.Second * 3)
	return a * b
}
func main() {
	//traditionalCase()
	contextCase()
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	<-ctx.Done()
}
