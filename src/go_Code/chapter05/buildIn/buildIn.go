package main

import "fmt"

func main() {

	// new用来分配内存，主要用来分配值类型，返回的是指针
	num1 := 100
	fmt.Printf("num1=%v,num1的类型是：%T,num1的地址是：%v\n", num1, num1, &num1)

	num2 := new(int) // *int
	*num2 = 200
	fmt.Printf("num2=%v,num2的类型是：%T,num2的地址是：%v\n", *num2, num2, &num2)

	// make主要为引用类型分配内存
}
