package main

import "fmt"

func calculate(num1 float64, num2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '*':
		res = num1 * num2
	case '/':
		res = num1 / num2
	default:
		fmt.Println("操作符有误")
	}
	return res
}

// 对函数返回值命名
func getSum(n1 float64, n2 float64) (sum float64) {
	sum = n1 + n2
	return
}

// 为函数取别名
type myFunType func(float64, float64) float64

// 将别名用在其他函数中
func function_variable(funcVar myFunType, number1 float64, number2 float64) float64 {
	return funcVar(number1, number2)
}

func sum(num1 float64, args ...float64) float64 {
	sum := num1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
func main() {
	//未完成某一功能的程序指令集合
	// 函数分为：自定义函数，系统函数
	// func 函数名 (形参列表) (返回值列表)  {
	// 	执行语句
	// 	return 返回值列表
	// }
	// 形参列表表示函数的输入，执行语句集合表示实现功能，函数可以有返回值也可以没有返回值
	var num1 float64 = 6.0
	var num2 float64 = 2.0
	var operator byte = '+'
	result := calculate(num1, num2, operator)
	fmt.Println(result)

	// go函数支持返回多个值，如果返回多个值，希望忽略某个返回值，使用 _ 表示占位忽略
	// 如果只有一个返回值，返回值列表不用写括号
	// 函数中的变量是局部的
	// 基本数据类型和数组默认都是值传递，即进行值拷贝，只在函数内修改并不会改变原函数的值
	// 如果希望函数内部的变量可以修改函数外的变量的值，可以传入变量的地址&，函数内以指针的方式操作变量(效果上看类似引用)
	// go函数不支持重载：函数名不变，参数列表不同
	// 函数也是一种数据类型，可以赋值给一个变量，该变量就是一个函数类型的变量，通过该变量可实现对函数的调用
	// 函数可以作为形参并且调用
	// go语言支持对函数返回值命名
	// go中支持可变参数:args名字自定义
	// 可变参数必须放在形参列表的最后面
	// ex：func sum(args...int) sum int{} 0到多个参数
	// ex：func sum(num int，args...int) sum int{} 1到多个参数
	// args本质上是slice切片，通过args[index]可以访问到各个值
	res1 := getSum(1, 2)
	fmt.Printf("res1=%v\n", res1)

	functionType := getSum
	fmt.Printf("result的数据类型是：%T\n", functionType)
	res2 := functionType(10, 20)
	fmt.Printf("res2=%v\n", res2)

	res3 := function_variable(getSum, 10, 10)
	fmt.Printf("res3=%v\n", res3)

	// 可变参数的使用
	res4 := sum(10, 100, 1000)
	fmt.Printf("res4=%v\n", res4)
}
