package main

import (
	"encoding/json"
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

type person struct {
	Name string
	Age  int
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

	// 创建结构体实例
	// 方式一：直接声明	var structDemo variableName
	// 方式二：var structDemo variableName = variableName{}
	person_02 := person{"mary", 20}
	person_02.Name = "tom"
	fmt.Println(person_02)
	// 方式三：var person *person = new(person)
	var person_03 *person = new(person)
	(*person_03).Name = "smith" //等价于 person_03.Name = "smith"
	(*person_03).Age = 25
	fmt.Println(*person_03)
	// 方式四：var person *person = &person{}
	var person_04 *person = &person{}
	(*person_04).Name = "scoot"
	// 等价于以下方式
	person_04.Name = "ashbur"
	fmt.Println(*person_04)
	var person_05 *person = &person{"mary", 30}
	fmt.Println(*person_05)

	// 结构体注意事项
	// 1.结构体字段在内存中都是连续的
	// 2.结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段（名字，个数，类型）
	// 3.结构体进行type重新定义（相当于取别名），编译器认为是新的数据类型，但是相互之间可以强制转换
	type student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	type stu student
	var stu_01 student
	var stu_02 stu = stu(stu_01)
	fmt.Println(stu_01, stu_02)
	// 4.struct每个字段上可以写一个tag，该tag可以通过反射机制获取，常见的使用场景是序列化和反序列化
	student_01 := student{"jones", 20}
	// 将student_01变量序列化为json格式字符串
	jsonstr, err := json.Marshal(student_01)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println(string(jsonstr))
}
