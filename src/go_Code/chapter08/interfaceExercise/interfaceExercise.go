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
	for i := 0; i < 10; i++ {
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
}
