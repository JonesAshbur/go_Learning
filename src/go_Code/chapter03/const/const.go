package main

import "fmt"

func main() {
	// 常量使用const修饰
	// 常量在定义时必须初始化
	// 常量不能修改
	// 常量只能修饰bool、数值类型(int、float、complex)、string类型
	// const (
	// 	a = 1
	// 	b = 2
	// )
	const (
		// iota是go语言的常量计数器，只能在常量的表达式中使用
		// iota在const关键字出现时将被重置为0
		// const中每新增一行常量声明将使iota计数一次
		num1 = iota // 0
		num2
		num3
	)
	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)
	fmt.Println("num3=", num3)
}
