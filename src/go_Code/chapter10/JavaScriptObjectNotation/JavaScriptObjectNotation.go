package main

import (
	"encoding/json"
	"fmt"
)

// 结构体序列化
type Person struct {
	Name     string `json:"name"` //假如前端想要小写，可以加tag标签，标签""内随意 （反射机制）
	Age      int
	Brithday string
	Salary   float64
	Skill    string
}

func testStruct() {
	person := Person{
		Name:     "ash",
		Age:      23,
		Brithday: "2002.1.1",
		Salary:   4567.8,
		Skill:    "coding",
	}
	data, error := json.Marshal(&person)
	if error != nil {
		fmt.Println("序列化失败：", error)
	}
	fmt.Printf("person序列化后结果：%v\n", string(data))
}

// map序列化
func testMap() {
	var mapDemo map[string]interface{} = make(map[string]interface{})
	mapDemo["name"] = "jones"
	mapDemo["age"] = "29"
	mapDemo["skill"] = "fight"
	data, error := json.Marshal(mapDemo)
	if error != nil {
		fmt.Println("序列化失败：", error)
	}
	fmt.Printf("mapDemo序列化后结果：%v\n", string(data))
}

// 切片序列化
func testSlice() {
	var slice []map[string]interface{}
	var map_01 map[string]interface{} = make(map[string]interface{})
	map_01["name"] = "jack"
	map_01["age"] = "20"
	map_01["skill"] = "eat"
	slice = append(slice, map_01)

	var map_02 map[string]interface{} = make(map[string]interface{})
	map_02["name"] = "jone"
	map_02["age"] = "25"
	map_02["skill"] = [2]string{"eat", "drink"}
	slice = append(slice, map_02)

	data, error := json.Marshal(slice)
	if error != nil {
		fmt.Println("序列化失败：", error)
	}
	fmt.Printf("slice序列化后结果：%v\n", string(data))
}

// 基本数据类型序列化
func testFloat64() {
	var num1 float64 = 10.11
	data, error := json.Marshal(num1)
	if error != nil {
		fmt.Println("序列化失败：", error)
	}
	fmt.Printf("num1序列化后结果：%v\n", string(data))
}
func main() {

	// JSON有利于机器解析和生成，有效提升网络传输效率，通常程序在网络传输时会先将数据（结构体，map等）序列化成JSON字符串，到接收方收到时，再进行反序列化恢复成原来的数据
	// JSON序列化是指，将key-value结构的数据类型序列化成json字符串
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
