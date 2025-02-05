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
}
