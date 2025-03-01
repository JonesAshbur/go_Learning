package main

import (
	"fmt"
)

type user struct {
	ID   int64
	Name string
	Age  uint8
}

type address struct {
	ID       int
	Province string
	City     string
}

// 泛型集合
type MapT[k comparable, V any] map[k]V

func (u user) String() string {
	return fmt.Sprintf("ID:%d,Name:%s,Age:%d", u.ID, u.Name, u.Age)
}
func (addr address) String() string {
	return fmt.Sprintf("ID:%d,Province:%s,City:%s", addr.ID, addr.Province, addr.City)
}

// 泛型接口
type GetKey[T comparable] interface {
	any
	GetKey() T
}

func (u user) GetKey() int64 {
	return u.ID
}
func (addr address) GetKey() int {
	return addr.ID
}

// 列表转集合
func ListToMap[k comparable, T GetKey[k]](list []T) map[k]T {
	mp := make(MapT[k, T], len(list))
	for _, v := range list {
		mp[v.GetKey()] = v
	}
	return mp
}
func InterfaceCase() {
	userList := []GetKey[int64]{
		user{ID: 1, Name: "jones", Age: 18},
		user{ID: 2, Name: "ash", Age: 28},
	}
	addressList := []GetKey[int]{
		address{ID: 1, Province: "直辖市", City: "北京"},
		address{ID: 2, Province: "直辖市", City: "天津"},
	}
	userMap := ListToMap[int64, GetKey[int64]](userList)
	fmt.Println(userMap)
	addressMap := ListToMap[int, GetKey[int]](addressList)
	fmt.Println(addressMap)
}
func main() {
	InterfaceCase()
}
