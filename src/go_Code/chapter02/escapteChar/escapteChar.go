package main

import "fmt" //主要包含格式化，输入，输出函数

func main() {

	fmt.Println("hello\nworld")

	fmt.Println("hello\tworld")

	fmt.Println("hello go\rworld") //不会换行，直接在此行从头覆写

	fmt.Println("hello\"world")

	fmt.Println("hello\\world")

	fmt.Println("姓名\t籍贯\t年龄\t住址\njacky\t河北\t23\t北京")

	fmt.Println("1111111111111111",
		"11111111111111111",
		"11111111111111111")
	//用","来进行换行
}

//	两种注释：
//	1.行注释：//
//	2.块注释：/*····*/	块注释不能嵌套
//	gofmt命令进行代码格式化并且输出
//	gofmt -w命令进行代码格式化并且重新写入
//	一行最多不超过80个字符

// fmt.Printf()用于做格式化输出
//  %+v包含字段名的值 %#vgo语法表示的值 %q单引号包起来 %08b指定二进制宽度为8，不足8位补0 %g紧凑输出
