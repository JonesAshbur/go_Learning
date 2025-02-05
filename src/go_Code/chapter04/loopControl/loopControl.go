package main

import (
	"fmt"
	"math/rand"
	"time"
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

	// 打印空心金字塔
	// 定义金字塔的层数
	layers := 5

	// 外层循环控制金字塔的层数
	for i := 1; i <= layers; i++ {
		// 打印空格，使金字塔居中
		for j := 1; j <= layers-i; j++ {
			fmt.Print(" ")
		}

		// 打印金字塔的左半边
		for k := 1; k <= 2*i-1; k++ {
			// 如果是第一层或最后一层，打印全部星号
			if i == 1 || i == layers {
				fmt.Print("*")
			} else {
				// 否则只打印第一个和最后一个星号，中间打印空格
				if k == 1 || k == 2*i-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}

		// 换行
		fmt.Println()
	}

	// 随机取数
	// 为了生成随机数需要用rand设置一个种子
	// go1.20开始弃用rand.Seed
	// time.Now().Unix()返回一个从1970:01:01的0时0分0秒到现在时间的秒数
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		num1 := rand.Intn(100) + 1 //	[0,100)
		fmt.Printf("num1=%d\n", num1)
		count++
		if num1 == 99 {
			break
		}
	}
	fmt.Printf("生成99一共使用了%d次\n", count)

	// break用于终止for循环或者跳出switch语句(最近的)
	// break出现在多层嵌套的语句块中时，可以通过标签指明要跳转到哪一层
	// lable1:
	// 	for i := 0; i < 4; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if j == 2 {
	// 				break lable1
	// 			}
	// 			fmt.Println("j=", j)
	// 		}
	// 	}

	// continue终止本次循环，继续执行下一次循环
	// continue出现在多层嵌套的语句块中时，可以通过标签指明要跳转到哪一层
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j=", j)
		}
	}

	// 跳转控制语句goto
	// 1.无条件转移到程序中指定的行
	// 2.通常与条件语句配合使用，可用来实现条件转移，跳出循环等功能
	// 3.不主张使用goto语句

	fmt.Println("1")
	goto label2
	fmt.Println("2")
	fmt.Println("3")
label2:
	fmt.Println("4")
	fmt.Println("5")

	// return使用在函数中，表示跳出所在的方法或者函数
}
