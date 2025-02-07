package main

import (
	"fmt"
	"strings"
)

// accumulator
func accumulator() func(int) int {
	var num1 int = 10
	return func(num2 int) int {
		num1 = num1 + num2
		return num1
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(filename string) string {
		if !strings.HasSuffix(filename, suffix) {
			return filename + suffix
		}
		return filename
	}
}

// defer
// 当执行到defer时，暂时不执行，将defer后面的语句压入defer栈
// 当函数执行完毕后，defer栈中的语句按照先入后出的顺序出栈，执行
func sum(n1 int, n2 int) int {
	defer fmt.Println("n1=", n1) //第三执行
	defer fmt.Println("n2=", n2) //第二执行
	res1 := n1 + n2
	fmt.Println("res1=", res1) //第一执行
	return res1
}
func main() {
	// 闭包：一个函数与其相关的引用环境组成的集合
	function_demo := accumulator()
	fmt.Println(function_demo(1))
	fmt.Println(function_demo(2))

	str1 := makeSuffix(".go")
	fmt.Printf("文件名处理后：%v\n", str1("closures.go"))

	// 在函数中需要经常创建资源，为了在函数执行完毕后即使释放资源，go提供了defer机制（延时机制）
	// defer将语句压入栈时，也会将相关的值拷贝同时入栈
	// defer后可以继续使用资源，当函数执行完毕后，会关闭资源
	res2 := sum(10, 20)
	fmt.Println("res2=", res2) //第四执行
}
