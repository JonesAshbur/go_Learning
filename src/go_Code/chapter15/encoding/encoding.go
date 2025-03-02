package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func EncodingCase() {
	type user struct {
		ID   int64
		Name string
		Age  uint8
	}
	u1 := user{
		ID:   1,
		Name: "ash",
		Age:  18,
	}
	//	序列化
	bytes1, err := json.Marshal(u1)
	fmt.Println(string(bytes1), err)
	//	反序列化
	u2 := user{}
	err = json.Unmarshal(bytes1, &u2)
	fmt.Println(u2, err)
	//	base64编码
	str1 := base64.StdEncoding.EncodeToString(bytes1)
	fmt.Println(str1)
	//	base64解码
	bytes2, err := base64.StdEncoding.DecodeString(str1)
	fmt.Println(string(bytes2), err)
	//	16进制编码
	str2 := hex.EncodeToString(bytes2)
	fmt.Println(str2)
	bytes3, err := hex.DecodeString(str2)
	fmt.Println(string(bytes3), err)
}

func main() {
	EncodingCase()
}
