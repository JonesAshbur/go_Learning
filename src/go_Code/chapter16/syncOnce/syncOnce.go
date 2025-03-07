// sync.Once作用：用来初始化单例资源

// sync.Once应用场景：
// 1.单例场景
// 2.仅加载一次的数据懒加载场景
package main

import (
	"fmt"
	"sync"
)

type onceMap struct {
	sync.Once
	data map[string]int
}

func (m *onceMap) LoadData() {
	m.Do(func() {
		list := []string{"A", "B", "C", "D"}
		for _, item := range list {
			_, ok := m.data[item]
			if !ok {
				m.data[item] = 0
			}
			m.data[item]++
		}
	})
}
func onceCase() {
	o := &onceMap{
		data: make(map[string]int),
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			o.LoadData()
		}()
	}
	wg.Wait()
	fmt.Println(o.data)
}
func main() {
	onceCase()
}
