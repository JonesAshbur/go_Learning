package main

import "fmt"

type USB interface {
	// 定义两个没有实现的方法
	Start()
	Stop()
}

// 手机
type Phone struct {
	Name string
}

// 相机
type Camera struct {
	Name string
}

// 让Phone实现USB接口的方法
func (mobile_phone Phone) Start() {
	fmt.Printf("%v手机开始工作\n", mobile_phone.Name)
}
func (mobile_phone Phone) Stop() {
	fmt.Printf("%v手机停止工作\n", mobile_phone.Name)
}

// phone独有方法
func (mobile_phone Phone) Call() {
	fmt.Printf("%v手机打电话\n", mobile_phone.Name)
}

// 相机实现USB接口方法
func (camera Camera) Start() {
	fmt.Printf("%v相机开始工作\n", camera.Name)
}
func (camera Camera) Stop() {
	fmt.Printf("%v相机停止工作\n", camera.Name)
}

type Computer struct {
}

func (computer Computer) Working(equipment USB) {
	equipment.Start()
	// 如果usb是指向phone结构体变量，则需要调用call方法
	// 类型断言
	if phone, flag := equipment.(Phone); flag {
		phone.Call()
	}
	equipment.Stop()
}

// 判断输入参数的类型
type Student struct {
	name string
	age  int
}

func TyprJudge(items ...interface{}) {

	for index, value := range items {
		// 如果不使用res变量接收判断结果，则在第79行需要再次进行类型断言（类型断言重复）
		switch res := value.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是：%v\n", index+1, res)
		case int, int8, int16, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是：%v\n", index+1, res)
		case float32, float64:
			fmt.Printf("第%v个参数是浮点类型，值是：%v\n", index+1, res)
		case string:
			fmt.Printf("第%v个参数是string类型，值是：%v\n", index+1, res)
		case byte:
			fmt.Printf("第%v个参数是byte类型，值是：%v\n", index+1, res)
		case Student:
			fmt.Printf("第%v个参数是struct类型，值是：%v\n", index+1, res)
		case *Student:
			// if student, flag := value.(*Student); flag {
			fmt.Printf("第%v个参数是struct指针类型，值是：%v\n", index+1, *res)
			// } else {
			// 	fmt.Printf("第%v个参数是struct指针类型，类型断言失败\n", index+1)
			// }

		default:
			fmt.Printf("第%v个参数类型不匹配\n", index+1)
		}
	}
}
func main() {

	var usbArr [3]USB
	usbArr[0] = Phone{"Apple"}
	usbArr[1] = Phone{"Xiaomi"}
	usbArr[2] = Camera{"nikon"}
	var computer Computer
	for _, value := range usbArr {
		computer.Working(value)
	}

	var n1 int16 = 1
	var n2 float32 = 1.1
	var n3 byte = '1'
	var n4 string = "1"
	var n5 bool = false
	n6 := "编程"

	student := Student{
		name: "jones",
		age:  18,
	}
	pupil := &Student{
		name: "ash",
		age:  20,
	}
	TyprJudge(n1, n2, n3, n4, n5, n6, student, pupil)
}
