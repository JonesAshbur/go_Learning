// sync.Con作用
// 1.设置一组协程根据条件阻塞，可以根据不同的条件阻塞
// 2。可以根据条件唤醒相应的协程

// sync.Con应用场景
// 应用于一发多收的场景，即一组协程需要等待某一个协程完成一些前置准备情况

// sync.Con注意事项
// 被叫方必须有锁
// 主叫方可以持有锁也可以不持有锁
// 尽可能减少无效唤醒
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func CondCase() {
	list := make([]int, 0)
	cond := sync.NewCond(&sync.Mutex{})
	go ReadList(&list, cond)
	go ReadList(&list, cond)
	go ReadList(&list, cond)
	time.Sleep(time.Second)
	InitList(&list, cond)
}
func InitList(list *[]int, c *sync.Cond) {
	//主叫方可以持有锁，也可以不持有锁
	c.L.Lock()
	defer c.L.Unlock()
	for i := 0; i < 10; i++ {
		*list = append(*list, i)
	}
	//唤醒所有条件等待的协程
	c.Broadcast()
}
func ReadList(list *[]int, c *sync.Cond) {
	//主叫方可以持有锁，也可以不持有锁
	c.L.Lock()
	defer c.L.Unlock()
	for len(*list) == 0 {
		fmt.Println("ReadList Wait")
		c.Wait()
	}
	fmt.Println("List数据：", *list)
}

type Queue struct {
	list []int
	cond *sync.Cond
}

func NewQueue() *Queue {
	q := &Queue{
		list: []int{},
		cond: sync.NewCond(&sync.Mutex{}),
	}
	return q
}
func (q *Queue) Push(item int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.list = append(q.list, item)
	//	当数据写入后，唤醒一个协程处理数据
	q.cond.Signal()
}
func (q *Queue) Pop(n int) []int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.list) < n {
		q.cond.Wait()
	}
	list := q.list[:n]
	q.list = q.list[n:]
	return list
}
func CondQueueCase() {
	q := NewQueue()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			list := q.Pop(n)
			fmt.Printf("%d:%d\n", n, list)
		}(i)
	}
	for i := 0; i < 100; i++ {
		q.Push(i)
	}
	wg.Wait()
}
func main() {
	//CondCase()
	CondQueueCase()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
