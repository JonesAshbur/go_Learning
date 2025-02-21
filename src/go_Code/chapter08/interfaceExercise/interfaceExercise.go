package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 定义一个hero结构体
type Hero struct {
	Name string
	Age  int
}

// 声明一个hero结构体类型切片
type HeroSlice []Hero

// 实现interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	// 等价于
	hs[i], hs[j] = hs[j], hs[i]
}
func main() {
	var heros HeroSlice
	for range 10 {
		hero := Hero{
			Name: fmt.Sprintf("英雄序号%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	// 排序前的顺序
	for _, value := range heros {
		fmt.Println(value)
	}
	// 调用sort包下的sort方法
	sort.Sort(heros)
	fmt.Println("排序后---------------------------")
	for _, value := range heros {
		fmt.Println(value)
	}

	// 接口和继承的比较
	// 1.当A结构体继承了B结构体，A就自动继承了B的字段和方法，并且可以直接使用
	// 2.当A结构体需要拓展某些功能，同时不希望破坏同B的继承关系，则可以去实现某个接口，可以认为接口是对继承机制的补充

	// 继承主要解决代码的复用性和可维护性
	// 接口主要设计好规范，让其他自定义类型去实现
	// 接口比继承更加灵活
	// 接口在一定程度上实现代码解耦
}
