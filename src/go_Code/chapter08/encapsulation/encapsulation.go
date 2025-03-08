package main

import (
	"fmt"

	"github.com/jonesashbur/golang_Learning/src/go_Code/chapter08/model"
)

func main() {

	// go语言仍然有面向对象语言的封装，继承和多态的特性，但是实现方式和其他OOP语言不同

	// 封装：将抽象出来的字段和字段的操作封装在一起，数据被保护在内部，程序的其他包只有通过被授权的操作（方法），才能对字段进行操作
	// 优点：隐藏实现细节，对数据进行验证，保证安全合理
	// 如何进行封装：对结构体中的属性进行封装，通过方法，包实现封装
	// 实现步骤：
	// 1.将结构体，字段的首字母小写
	// 2.给结构体所在包提供一个工厂模式的函数，首字母大写
	// 3.提供一个首字母大写的Set方法，用于对属性判断并且赋值
	// func (var structVariable) Setfun(参数列表) (返回值列表){
	// 数据验证逻辑
	// var.字段 = 参数
	// }
	// 4.提供一个首字母大写的Get方法，用于获取属性的值
	// func (var 结构体类型名) Getxxx(){
	// return var.字段
	// }
	// go本身对面向对象特性做了简化，go开发中并没有特别强调封装

	p := model.Newperson("jones")
	p.SetAge(26)
	p.SetSalary(3456.7)
	fmt.Println("name:", (*p).Name)
	fmt.Println("age:", (*p).GetAge())
	fmt.Println("salary:", (*p).GetSalary())
	fmt.Println(*p)

}
