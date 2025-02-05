package main

import (
	"fmt"
)

func main() {

	// ex1:
	// for cycleTimes1 := 1; cycleTimes1 <= 5; cycleTimes1++ {
	// 	fmt.Println("hello golang")
	// }

	// for循环注意事项：
	// 1.循环条件是一个返回布尔值的表达式
	// 2.循环变量的初始化可以写在for循环的外面
	// 3.for ;; {}无限循环，配合break使用
	// 4.for-range遍历字符串和数组
	// ex2:
	// cycleTimes2 := 1
	// for cycleTimes2 <= 5 {
	// 	fmt.Println("hello golang")
	// 	cycleTimes2++
	// }
	// 字符串遍历方式1：默认按照字节遍历，一个汉字在utf-8编码中占3个字节，遍历中文会出乱码
	// 需要将字符串转成切片
	var str1 string = "hello,golang中文"
	for str_loop := 0; str_loop < len(str1); str_loop++ {
		fmt.Printf("%c\n", str1[str_loop])
	}
	str2 := []rune(str1)
	for str_loop := 0; str_loop < len(str2); str_loop++ {
		fmt.Printf("%c\n", str2[str_loop])
	}
	// 字符串遍历方式2：可以遍历中文
	for index, val := range str1 {
		fmt.Printf("inedx=%d,val=%c\n", index, val)
	}

	//go语言本身没有while循环和do while循环

	// for循环实现while循环效果
	// 循环变量初始化
	// for {
	// 	if 循环条件表达式 {
	// 		break //跳出循环
	// 	}
	// 	循环语句
	// 	循环变量迭代
	// }

	// for循环实现do while
	// 循环变量初始化
	// for {
	// 	循环语句
	// 	循环变量迭代
	// 	if 条件表达式 {
	// 		break
	// 	}
	// }
}
