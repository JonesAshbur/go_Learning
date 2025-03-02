package main

import "fmt"

// 泛型结构体
type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}

// 泛型receiver
func (mystruct MyStruct[T]) GetData() T {
	return mystruct.Data
}
func ReceiverCase() {
	data := 18
	mystruct1 := MyStruct[*int]{
		Name: "ash",
		Data: &data,
	}
	str := "hello golang"
	mystruct2 := MyStruct[*string]{
		Name: "ash",
		Data: &str,
	}
	dataContent1 := mystruct1.GetData()
	dataContent2 := mystruct2.GetData()
	fmt.Println(*dataContent1)
	fmt.Println(*dataContent2)
	//fmt.Println(mystruct1, mystruct2)
}
func main() {
	ReceiverCase()
}
