package main

import "fmt"

// 传统写法
func getMaxNumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func getMaxNumFloat(a, b float64) float64 {
//	if a > b {
//		return a
//	}
//	return b
//}

// 泛型写法
// any泛型类型的类型约束
type customerT interface {
	//单行取并集，多行取交集
	int | int16 | int64 | float32 | float64 | ~int32 //包含int32及其所有衍生类型
}

// type Myint32 int32   //int32衍生类型，是一个新类型
// type Myint64 = int64 //别名
func getMaxNum[T customerT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 内置类型comparable
func getBuildInComparable[T comparable](a, b T) bool {
	//comparable 类型 只支持 == != 两个操作
	//	支持所有可用于== !=的基本数据类型
	if a == b {
		return true
	}
	return false
}

// 内置类型any
func getBuildInAny[T any](a, b T) {
	//comparable 类型 只支持 == != 两个操作
	//	支持所有可用于== !=的基本数据类型
	fmt.Println("内置any类型约束：", a, b)
}
func main() {

	// 泛型即开发过程中编写适用于所有类型的模板，只有在使用的时候才能确定其真正的类型
	// 增加代码复用，从同类型到不同类型的代码复用
	// 应用于不同类型间代码复用的场景，即不需要写相同的处理逻辑时，最适合用泛型
	// 增加了编译器负担，降低了编译效率
	// 匿名结构体和匿名函数不支持泛型的定义
	// 不支持类型断言
	// 不支持泛型方法，只能通过receiver来实现方法的泛型处理
	// ~后的类型必须是基本数据类型，不能是接口类型

	var a int = 1
	var b int = 2

	var c float64 = 3
	var d float64 = 4
	//	不使用泛型数字比较
	res01 := getMaxNumInt(a, b)
	fmt.Println("不使用泛型数字比较结果：", res01)
	//显示指定泛型类型
	res02 := getMaxNum[float64](c, d)
	fmt.Println("使用泛型数字比较结果：", res02)

	var str1 string = "abc"
	var str2 string = "def"
	res03 := getBuildInComparable(str1, str2)
	fmt.Println("comparable比较结果：", res03)

	getBuildInAny(str1, str2)
}
