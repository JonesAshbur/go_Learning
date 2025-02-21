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
	fmt.Println("手机开始工作")
}
func (mobile_phone Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 相机实现USB接口方法
func (camera Camera) Start() {
	fmt.Println("相机开始工作")
}
func (camera Camera) Stop() {
	fmt.Println("相机停止工作")
}

// 计算机
type Computer struct {
}

// 定义一个方法接收USB接口类型的变量（实现了usb接口里面的所有方法 ）
func (pc Computer) Working(equipment USB) {
	equipment.Start()
	equipment.Stop()
}
func main() {
	var phone_01 Phone
	var interface_01 USB = phone_01
	interface_01.Start()
	interface_01.Stop()
	// 变量（实例）具有多种形态，OOP的第三大特征，
	// 多态特征在go语言中通过接口实现，可以通过统一的接口调用不同的实现，这时接口变量就呈现了不同的形态

	// 接口体现多态特征
	// 1.多态参数：接收的参数可以是不同的实例，例如USB接口
	// 2.多态数组

	// 定义一个USB接口数组(数组只能存放相同数据类型的数据，但是借助接口可以实现多态数组)
	var usbArr [3]USB
	usbArr[0] = Phone{"Apple"}
	usbArr[1] = Phone{"Xiaomi"}
	usbArr[2] = Camera{"nikon"}
	fmt.Println(usbArr)

}
