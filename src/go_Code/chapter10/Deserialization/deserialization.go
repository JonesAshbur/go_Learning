package main

import (
	"encoding/json"
	"fmt"
)

// JSON反序列化struct
type Person struct {
	Name     string
	Age      int
	Brithday string
	Salary   float64
	Skill    string
}

func unmarshalStruct() {
	str_01 := "{\"name\":\"ash\",\"Age\":23,\"Brithday\":\"2002.1.1\",\"Salary\":4567.8,\"Skill\":\"coding\"}"
	var person Person
	error := json.Unmarshal([]byte(str_01), &person)
	if error != nil {
		fmt.Println("反序列化失败：", error)
	}
	fmt.Println("person:", person)
}

// JSON反序列化map
func unmarshalMap() {
	str_02 := "{\"age\":\"29\",\"name\":\"jones\",\"skill\":\"fight\"}"
	var mapDemo map[string]interface{} //反序列化不需要make,make已经被封装到unmarshal中
	error := json.Unmarshal([]byte(str_02), &mapDemo)
	if error != nil {
		fmt.Println("反序列化失败：", error)
	}
	fmt.Println("mapDemo:", mapDemo)
}

// slice反序列化
func unmarshalSlice() {
	str_02 := "[{\"age\":\"20\",\"name\":\"jack\",\"skill\":\"eat\"}," +
		"{\"age\":\"25\",\"name\":\"jone\",\"skill\":[\"eat\",\"drink\"]}]"
	var slice []map[string]interface{}
	error := json.Unmarshal([]byte(str_02), &slice)
	if error != nil {
		fmt.Println("反序列化失败：", error)
	}
	fmt.Println("slice:", slice)
}
func main() {

	// 反序列化
	// 将JSON转换成对应的数据类型
	// 反序列化一个JSON数据时，要确保反序列化后的数据与序列化前的数据类型保持一致
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
