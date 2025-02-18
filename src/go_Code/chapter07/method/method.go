package main

import "fmt"

type example struct {
	Num1 int
}

// 名为testFunc的函数是example的一个方法
func (a example) testFunc() {
	a.Num1 = 5
	fmt.Println("a.Num1=", a.Num1)
}

func (c example) calculate1(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func (b *example) calculate2() {
	fmt.Printf("b的地址：%p\n", b)
}

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}
func (i *integer) add() {
	*i++
}

type student struct {
	Name string
	Age  int
}

func (stu *student) String() string {
	str := fmt.Sprintf("Name=[%v],Age=[%v]", stu.Name, stu.Age)
	return str
}
func main() {

	// golang中的方法是作用在指定的数据类型上（和指定的数据类型绑定），因此自定义数据类型，都可以有方法
	// 方法和指定数据类型绑定，只能通过指定数据类型的变量来调用，不能被其他数据类型的变量调用
	// func (a example) testFunc() 其中a表示哪个example变量调用，a就是它的副本
	// 方法的调用和传参机制基本和函数一样，不同的地方是方法调用时，会将调用方法的变量一并当作实参传递给方法
	// 为了提高效率，通常方法和结构体的指针类型绑定
	// 如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个类型的String()进行输出
	var a example
	a.Num1 = 10
	fmt.Println("a.Num1=", a.Num1)

	a.testFunc()

	var n int = 1000
	result := a.calculate1(n)
	fmt.Println("result=", result)

	var b example
	(&b).calculate2()
	fmt.Printf("b的地址：%p\n", &b)

	var i integer = 10
	i.print()
	i.add()
	fmt.Println("i=", i)

	demo := student{
		Name: "jones",
		Age:  28,
	}
	res := demo.String()
	fmt.Println(res)
	fmt.Println(&demo)

	// 方法和函数的区别
	// 1.调用方式不一样：
	// 函数调用方式：函数名（实参列表）
	// 方法调用方式：变量名.方法名（实参列表）
	// 2，对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
	// 3.对于方法，接收者为值类型时，可以直接用指针类型的变量调用方法，反过来也可以

	// 创建结构体变量时指定字段的值：
	// 方式一：创建时直接指定字段的值 var demo structVariable = {"xxx",x,x...}
	// 方式二：创建结构体字段时，将字段名和字段值写在一起
	// demo := structVariable {
	// 	name : "xxx",
	// 	age = 20,
	// }
	// 方式三：返回结构体的指针类型
	// demo := &structVariable {
	// 	name : "xxx",
	// 	age = 20,
	// }
}
