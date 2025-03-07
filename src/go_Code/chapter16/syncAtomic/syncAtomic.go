// sync/atomic作用
// atomic提供底层的原子级内存操作，用于实现同步算法
// 应用场景：通过内存实现通信
// 注意事项：
// 1.atomic属于底层原子级操作，无必要条件下使用channel或sync包其他函数实现同步算法
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCase1() {
	var num int64 = 5
	atomic.StoreInt64(&num, 10)
	fmt.Println("原数据的值：", atomic.LoadInt64(&num))
	fmt.Println("原数据基础上累加后的值：", atomic.AddInt64(&num, 10))
	fmt.Println("交换变量的值，并且返回原有的值：", atomic.SwapInt64(&num, 30))
	fmt.Println("交换后的值：", atomic.LoadInt64(&num))
	fmt.Println("比较并替换原有的值，并返回是否替换成功：", atomic.CompareAndSwapInt64(&num, 30, 40))
	fmt.Println("比较并且替换后的值：", atomic.LoadInt64(&num))
	fmt.Println("比较并替换原有的值，并返回是否替换成功：", atomic.CompareAndSwapInt64(&num, 30, 50))
	fmt.Println("比较并且替换后的值：", atomic.LoadInt64(&num))
}

type atomicCounter struct {
	count int64
}

// 通过锁实现原子级操作(上锁和解锁之间的操作是原子性)
func AtomicCase2() {
	var count int64
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			count++ //lock之间是原子操作
			lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func (c *atomicCounter) Add() {
	atomic.AddInt64(&c.count, 1)
}
func (c *atomicCounter) Load() int64 {
	return atomic.LoadInt64(&c.count)
}

// 通过atomicCount计数
func AtomicCase3() {
	var count = atomicCounter{}

	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Add()
		}()
	}
	wg.Wait()
	fmt.Println(count.Load())
}

func AtomicCase4() {
	list := []string{"a", "b", "c", "d", "e", "f", "g"}
	//定义一个原子值
	var atomicMp atomic.Value
	//	定义一个集合
	mp := map[string]int{}
	//集合不能用于对比，所以传入地址
	atomicMp.Store(&mp)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		atomicLable:
			m := atomicMp.Load().(*map[string]int)
			newMp := map[string]int{}
			for k, v := range *m {
				newMp[k] = v
			}
			for _, item := range list {
				_, ok := newMp[item]
				if !ok {
					newMp[item] = 0
				}
				newMp[item]++
			}
			swap := atomicMp.CompareAndSwap(m, &newMp)
			if !swap {
				goto atomicLable
			}
			//fmt.Println(swap)
		}()
	}
	wg.Wait()
	fmt.Println(atomicMp.Load())
}
func main() {
	//AtomicCase1()
	//AtomicCase2()
	//AtomicCase3()
	AtomicCase4()
}
