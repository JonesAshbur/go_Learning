package main

import (
	"fmt"
	"sort"
)

type SortUser struct {
	ID   int64
	Name string
	Age  uint8
}

type ByID []SortUser

// 获取长度
func (user ByID) Len() int {
	return len(user)
}

// 交换位置
func (user ByID) Swap(i, j int) {
	user[i], user[j] = user[j], user[i]
}

// 返回i位置的ID是否小于j位置的ID
func (user ByID) Less(i, j int) bool {
	return user[i].ID < user[j].ID
}
func SortCase() {
	list := []SortUser{
		{ID: 1, Name: "ash", Age: 68},
		{ID: 4, Name: "jack", Age: 28},
		{ID: 7, Name: "jones", Age: 48},
		{ID: 2, Name: "mary", Age: 32},
		{ID: 8, Name: "ashbur", Age: 48},
		{ID: 3, Name: "alfred", Age: 28},
		{ID: 5, Name: "smith", Age: 32},
		{ID: 9, Name: "ken", Age: 78},
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Age < list[j].Age
	})
	fmt.Println(list)
	//	实现sort.Interface()接口
	sort.Sort(ByID(list))
	fmt.Println(list)
}
func main() {
	SortCase()
}
