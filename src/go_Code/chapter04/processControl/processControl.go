package main

import (
	"fmt"
)

func main() {

	// 分支控制：单分支，双分支，多分支
	// 条件判断必须是条件表达式，不能是赋值语句
	// fmt.Println("今天学习了吗？(输入true或者false)")
	// var isStudy bool
	// fmt.Scanln(&isStudy)
	// if isStudy {
	// 	fmt.Println("good job")
	// } else {
	// 	fmt.Println("bad job")
	// }

	// ex:golang支持在if中直接定义一个变量
	if num1 := 10; num1 > 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	//多分支控制:只会有一个符合条件的输出
	if num2 := 10; num2 > 20 {
		fmt.Println("yes")
	} else if num2 > 5 {
		fmt.Println("ok")
	} else if num2 > 7 {
		fmt.Println("no")
	} else { // else可以省略
		fmt.Println("overflow")
	}

	// switch分支结构
	// 匹配项后面不需要加break（默认会有）
	// golang switch分支结构中case后面的表达式可以有多个，使用逗号间隔
	// 默认情况下，执行完case语句块就退出switch控制结构
	// 注意事项：
	// 1.case后面是一个表达式（常量，变量，有返回值的函数）
	// 2.case后面表达式的值的数据类型必须和switch后面的表达式数据类型保持一致
	// 3.case后面的表达式如果是一个字面量（常量），则不能重复
	// 4.default语句不是必须的
	// 5.switch后面可以不加表达式，类似于if-else分支使用
	// 6.switch后面可以直接声明/定义一个变量，以分号结束（不推荐）
	// 7.switch穿透（fallthrough），如果在case语句后加fallthrough，则会继续执行下一个case语句
	// 8.Type switch：switch语句可被用于Type switch来判断interface变量中实际指向的变量类型
	var key1 int8 = 5
	switch key1 {
	case 1, 2:
		fmt.Println("不是我")
	case 3, 4:
		fmt.Println("也不是我")
	case 5, 6:
		fmt.Println("是我")
	default:
		fmt.Println("找不到")
	}

	// 当作if-else使用
	var key2 int8 = 5
	switch {
	case key2 == 5:
		fmt.Println("匹配")
	default:
		fmt.Println("找不到")
	}

	switch key3 := 5; {
	case key3 == 5:
		fmt.Println("yes")
	default:
		fmt.Println("no")
	}

	var key4 int8 = 5
	switch {
	case key4 == 5:
		fmt.Println("第一次匹配")
		fallthrough //	穿透
	case key4 == 5:
		fmt.Println("第二次匹配（穿透）")
	default:
		fmt.Println("找不到")
	}

	//Type switch
	var key5 interface{}
	var key6 = 1.0
	key5 = key6
	switch variable := key5.(type) {
	case int8:
		fmt.Printf("variable是%T类型", variable)
	case float64:
		fmt.Printf("variable是%T类型", variable)
	default:
		fmt.Println("找不到该数据类型")
	}

	// switch和if-else的比较
	// 如果判断的具体数据不多，并且符合基本以下数据类型（整数，浮点数，字符串，字符），建议switch，简洁高效
	// 如果对区间判断，且结果为bool类型则使用if-else，范围更广
}
