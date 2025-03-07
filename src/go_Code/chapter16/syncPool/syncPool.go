// sync.Pool作用：
//创建一个临时对象池，缓存一组对象用于重复利用，以此来减少内存分配和GC压力

// sync.Pool应用场景：
// 可用于连接池（grpc客户端，网络连接）等场景

// sync.Pool注意事项
// 1.用于缓存一些创建成本较高，使用比较频繁的对象
// 2.Pool的长度默认是CPU线程数
// 3.存储在Pool中的对象随时都可能在不被通知的情况下被回收
// 4.没有一定创建成本的对象不建议使用对象池
package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

func PoolCase() {
	target := "192.168.23.15"
	pool, err := GetPool(target)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		conn := &Conn{
			ID:     int64(i),
			Target: target,
			Status: ON,
		}
		pool.Put(conn)
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				conn := pool.Get()
				fmt.Println(conn.ID)
				pool.Put(conn)
			}
		}()
	}
	wg.Wait()
}

const (
	ON  = 1
	OFF = 0
)

type Conn struct {
	ID     int64
	Target string
	Status int
}

func NewConn(target string) *Conn {
	return &Conn{
		ID:     rand.Int63(),
		Target: target,
		Status: ON,
	}
}
func (c *Conn) GetStatus() int {
	return c.Status
}

type ConnPool struct {
	sync.Pool
}

func GetPool(target string) (*ConnPool, error) {
	return &ConnPool{
		Pool: sync.Pool{
			New: func() any {
				return NewConn(target)
			},
		},
	}, nil
}
func (c *ConnPool) Get() *Conn {
	conn := c.Pool.Get().(*Conn)
	if conn.GetStatus() == OFF {
		conn = c.Pool.New().(*Conn)
	}
	return conn
}
func (c *ConnPool) Put(conn *Conn) {
	if conn.GetStatus() == OFF {
		return
	}
	c.Pool.Put(conn)
}
func main() {
	PoolCase()
}
