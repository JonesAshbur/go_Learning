// Mutex,RWMutex作用
// 并发场景下通过锁机制解决数据竞争问题

// 应用场景
// 协程安全，数据竞争

// 注意事项
// 避免使用锁
// 合理使用，不能滥用

package main

import (
	"fmt"
	"sync"
)

func MutexCase() {

}

// SingleRoutine 单协程，协程安全
func SingleRoutine() {
	mp := make(map[string]int, 0)
	list := []string{"A", "B", "C", "D"}
	for i := 0; i < 20; i++ {
		for _, item := range list {
			_, ok := mp[item]
			if !ok {
				mp[item] = 0
			}
			mp[item] += 1
		}
	}
	fmt.Println(mp)
}

// MultipleRoutine 多协程，非协程安全
func MultipleRoutine() {
	mp := make(map[string]int, 0)
	list := []string{"A", "B", "C", "D"}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, item := range list {
				_, ok := mp[item]
				if !ok {
					mp[item] = 0
				}
				mp[item] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp)
}

// MultipleSafeRoutine 多协程，互斥锁，协程安全
func MultipleSafeRoutine() {
	type safeMap struct {
		data map[string]int
		sync.Mutex
	}
	mp := safeMap{
		data:  make(map[string]int, 0),
		Mutex: sync.Mutex{},
	}
	list := []string{"A", "B", "C", "D"}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mp.Lock()
			defer mp.Unlock()
			for _, item := range list {
				_, ok := mp.data[item]
				if !ok {
					mp.data[item] = 0
				}
				mp.data[item] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp.data)
}

type Cache struct {
	data map[string]string
	sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data:    make(map[string]string, 0),
		RWMutex: sync.RWMutex{},
	}
}
func (c *Cache) Get(key string) string {
	c.RLock()
	defer c.RUnlock()
	value, ok := c.data[key]
	if ok {
		return value
	}
	return ""
}
func (c *Cache) Set(key, value string) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}

// MultipleSafeRoutineByRWMutex 读写锁，读写，写写加锁，读读不加锁
func MultipleSafeRoutineByRWMutex() {
	c := NewCache()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Set("name", "ash")
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Get("name")
			fmt.Println(c.Get("name"))
		}()
	}

}
func main() {
	//SingleRoutine()
	//MultipleRoutine()
	//MultipleSafeRoutine()
	MultipleSafeRoutineByRWMutex()
}
