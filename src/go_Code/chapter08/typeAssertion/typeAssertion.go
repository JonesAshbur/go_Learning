package main

import (
	"fmt"
)

type Point struct {
	x1 int
	y1 int
}

func main() {

	// 类型断言：由于接口是一般类型，不知道具体类型，如果要转成具体类型，就要使用类型断言
	// 空接口重新转换成对应类型
	var a1 interface{} //空接口
	var point Point = Point{1, 2}
	a1 = point
	var b1 Point = a1.(Point) //表示判断a是否是指向point类型的变量
	// b1 = a1 报错
	fmt.Println(b1)

	var a2 interface{}
	var b2 float64 = 1.2
	a2 = b2
	temp1 := a2.(float64) //要明确原先的接口指向的是什么类型
	fmt.Printf("temp的数据类型是：%T，值为：%v\n", temp1, temp1)

	// 类型断言时附带检测机制（不报panic）
	var a3 interface{}
	var b3 float64 = 1.2
	a3 = b3
	if temp2, flag := a3.(float32); flag {
		fmt.Println("转换成功")
		fmt.Printf("temp的数据类型是：%T，值为：%v\n", temp2, temp2)
	} else {
		fmt.Println("转换失败")
	}
	fmt.Println("后续代码......")
}
