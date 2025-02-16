package main

import (
	"fmt"
)

type example struct {
	Name    string
	Age     int
	Score   [5]float32
	ptr     *int
	slice   []int
	mapDemo map[int]int
}

func main() {
	// go语言支持面向对象的编程特性，和传统面向对象编程语言有区别，并不是纯粹的面向对象编程语言
	// go基于struct实现OOP特性
	// go去掉了传统OOP语言的继承，方法重载，构造函数和析构函数，隐藏的this指针等
	// go仍然有封装，继承，多态的特性，但是实现形式和传统OOP不同，比如继承，go没有extends关键字，继承通过匿名字段来实现
	// 通过接口interface关联，耦合性较低，非常灵活，面向接口编程是go中非常重要的特性

	// 结构体和结构体变量（实例）之间的关系：
	// 结构体是自定义的数据类型
	// 结构体变量（实例）是具体的，实际的代表一个具体变量
	// 声明结构体
	// type structDemo struct{
	// 	demo_1 type
	// 	demo_2 type
	// 	...
	// }
	// 结构体变量名字和字段名字大小写与其作用域相对应
	// 字段的类型可以是值类型，也可以是引用类型
	// 如果字段没有赋值，默认是零值，数组的默认值与其元素的类型有关，指针，slice，map的默认值是nil，即没有分配空间

	var person_01 example
	fmt.Println(person_01)
	// 验证默认nil
	if person_01.ptr == nil {
		fmt.Println("ptr Success")
	}
	if person_01.slice == nil {
		fmt.Println("slice Success")
	}
	if person_01.mapDemo == nil {
		fmt.Println("mapDemo Success")
	}

	person_01.slice = make([]int, 5)
	person_01.slice[0] = 10
	fmt.Println(person_01.slice)

	person_01.mapDemo = make(map[int]int, 2)
	person_01.mapDemo[0] = 1
	person_01.mapDemo[1] = 2
	fmt.Println(person_01.mapDemo)

}
