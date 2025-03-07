// sync.Map 线程安全集合，内部通过原子访问和锁机制实现结合的线程安全

// 适合读多写少的应用场景，在key存在的情况下可以无锁修改value
package main

import (
	"fmt"
	"sync"
)

func SyncMapCase() {
	mp := sync.Map{}
	//设置键值对
	mp.Store("name", "ash")
	mp.Store("e-mail", "jonesashbur@xxx.com")
	//通过key获取value，不存在返回nil，ok则返回false
	fmt.Println(mp.Load("name"))
	fmt.Println(mp.Load("e-mail"))
	//通过key获取value，如果不存在则设置指定value值并且返回
	//ok为ture表示key存在并且返回其值，为false表示key不存在设置后返回
	fmt.Println(mp.LoadOrStore("hobby", "football"))
	fmt.Println(mp.LoadOrStore("hobby", "football"))
	//存在，值依旧是football
	fmt.Println(mp.LoadOrStore("hobby", "basketball"))
	//根据key获取value，并且删除该key
	//ok为true表示key存在，为false表示key不存在
	fmt.Println(mp.LoadAndDelete("hobby"))
	fmt.Println(mp.LoadAndDelete("hobby"))
	//	为集合设置迭代函数，将为集合中的每一个键值对顺序调用该函数，如果返回false则停止迭代
	mp.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}
func main() {
	SyncMapCase()
}
