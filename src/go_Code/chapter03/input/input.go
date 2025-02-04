package main

import "fmt"

func main() {

	// fmt.Scanln()
	// var name1 string
	// var age1 byte
	// var score1 float32
	// var ismarried1 bool
	// fmt.Println("请输入姓名:")
	// fmt.Scanln(&name1)
	// fmt.Println("请输入年龄:")
	// fmt.Scanln(&age1)
	// fmt.Println("请输入分数:")
	// fmt.Scanln(&score1)
	// fmt.Println("请输入是否结婚:")
	// fmt.Scanln(&ismarried1)
	// fmt.Printf("姓名：%v\n年龄：%v\n分数：%v\n是否已婚：%v\n", name1, age1, score1, ismarried1)

	// fmt.Scanf()
	// var name2 string
	// var age2 byte
	// var score2 float32
	// var ismarried2 bool
	// fmt.Println("请输入姓名，年龄，分数，是否已婚，使用空格隔开")
	// fmt.Scanf("%s %d %f %t", &name2, &age2, &score2, &ismarried2)
	// fmt.Printf("姓名：%v\n年龄：%v\n分数：%v\n是否已婚：%v\n", name2, age2, score2, ismarried2)

	// 在golang中不能直接使用二进制来表示一个整数，八进制和十六进制可以直接输出
	var num1 byte = 5
	fmt.Printf("5对应的二进制数是%b\n", num1)

	var num2 byte = 011
	fmt.Println("num2=", num2)

	var num3 byte = 0x11
	fmt.Println("num3=", num3)

}
