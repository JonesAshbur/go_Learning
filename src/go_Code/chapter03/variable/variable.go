package main

import "fmt"

func main() {

	//声明变量
	var num1 int

	//变量赋值
	num1 = 10

	//使用变量
	fmt.Println("num1=", num1)
	//var num1 int = 10

	//变量使用的注意事项：
	// 1.变量表示内存中的一片存储区域
	// 2.变量名 + 变量数据类型
	// 3.变量声明后如果不赋值则使用默认值，int的默认值是0，string默认为空串，小数默认为0
	var num2 int
	fmt.Println("num2=", num2)
	// 4.根据值自行判断变量的数据类型（类型推导）声明时赋值可省略数据类型
	var num3 = 10
	fmt.Println("num3=", num3)
	// 省略var，:=左边的变量不能是已经声明过的
	name1 := "Tom"
	//等价于 var name1 string = "Tom"
	fmt.Println("name1:", name1)
	// 5.多变量声明，支持一次性声明多个变量
	var num4, num5, num6 int
	fmt.Println("num4=", num4, "num5=", num5, "num6=", num6)
	// 6.一次性声明多个不同数据类型的变量
	var name2, age1, gender1 = "Tom", 23, "male"
	fmt.Println("name2:", name2, "age1:", age1, "gender1:", gender1)
	// 类型推导
	name3, age2, gender2 := "jacky", 24, "fmale"
	fmt.Println("name3:", name3, "age2:", age2, "gender2:", gender2)
	// 7.该区域的数据值可以在同一类型范围内不断变化(不能改变变量的作用域和数据类型)
	var num7 int = 100
	num7 = 200
	fmt.Println("num7=", num7)
	// 8.变量在同一个作用域内不能重名
	// 9.变量的三要素：变量名，值，数据类型
	// 全局变量
	fmt.Println("name4:", name4, "age3:", age3, "score1:", score1)
	fmt.Println("name5:", name5, "age4:", age4, "score2:", score2)
}

// 全局变量的定义方式1
var name4 string = "jones"
var age3 int = 25
var score1 float32 = 100.0

//	全局变量的定义方式2
var (
	name5  = "ashbur"
	age4   = 30
	score2 = 90
)
