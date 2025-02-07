package main

import "fmt"

func init() {
	fmt.Println("这是init函数")
}

// 全局匿名函数
var (
	AnonymousFunction = func(n1 int, n2 int) int {
		return n1 + n2
	}
)

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

	// 每一个源文件都可以包含一个init函数，该函数会在main函数执行之前，被go运行框架调用
	// 在init函数中完成初始化工作
	// 如果一个文件包含全局变量定义，init函数和main函数，执行顺序是全局变量定义，init函数，main函数

	// 匿名函数
	// 如果某个函数只希望执行一次，可以使用匿名函数，匿名函数也可以实现多次调用
	// 使用方式：1.定义匿名函数时直接调用（这种方式匿名函数只能调用一次）	2.将匿名函数赋值给一个变量
	// 如果将匿名函数复制给一个全局变量，那么这个匿名函数就成为全局匿名函数
	res5 := func(n1 int, n2 int) int {
		return n1 + n2
	}(100, 200)
	fmt.Printf("res5=%v\n", res5)

	res6 := func(n1 int, n2 int) int {
		return n1 - n2
	}
	// res6可反复调用，但是res6不是函数名
	res7 := res6(1000, 100)
	fmt.Printf("res7=%v\n", res7)

	// 调用全局匿名函数
	res8 := AnonymousFunction(100, 900)
	fmt.Printf("res8=%v\n", res8)

	// 函数参数的传递：
	// 1.值传递和引用传递
	// 2.不管是值传递还是引用传递，传递给函数的都是变量的副本，值传递是值拷贝，引用传递是地址拷贝
	// 3.地址拷贝的效率更高，因为数据量小
	// 4.值拷贝数据越大，效率越低
}
