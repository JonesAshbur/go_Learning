package main

import (
	"errors"
	"fmt"
)

func test1() {
	defer func() {

		if error := recover(); error != nil {
			fmt.Println("error=", error)
			fmt.Println("做出提醒")
		}
	}()
	num1 := 1
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func readConfig(name string) (err error) {
	if name == "config.name" {
		return nil
	} else {
		return errors.New("读取config文件失败")
	}
}

func test2() {
	error := readConfig("config.name1")
	if error != nil {
		panic(error)
	}
	fmt.Println("test2剩余代码继续执行")
}
func main() {

	// 测试
	test1()
	fmt.Println("继续执行test1()后面的代码")

	//	默认情况下出现错误（panic），直接退出

	// 错误处理方式：defer，panic，recover
	// 错误处理的好处：程序不会轻易退出，加入预警代码，程序更加健壮
	// go编译器抛出panic，在defer中recover捕获这个异常
	// go也支持自定义错误，使用函数error.New和panic
	// error.New()返回一个error类型的值，表示一个错误
	// panic内置函数，接受一个interface{}类型的值(可以是任何值)作为参数，可以接受error类型的变量，输出错误信息并且退出程序

	// 测试自定义错误
	test2()
	fmt.Println("继续执行test2()后面的代码")
}
