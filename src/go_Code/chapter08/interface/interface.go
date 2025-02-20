package main

import "fmt"

type USB interface {
	// 定义两个没有实现的方法
	Start()
	Stop()
}

// 手机
type Phone struct {
}

// 相机
type Camera struct {
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

	// 创建结构体变量
	// computer := Computer{}
	// phone := Phone{}
	// camera := Camera{}

	// computer.Working(phone)
	// computer.Working(camera)
	var phone_01 Phone
	var interface_01 USB = phone_01
	interface_01.Start()
	interface_01.Stop()
	// 接口：interface类型可以定义一组方法，但是这些方法不需要实现，并且interface不能包含任何变量，到某个自定义类型要使用的时候，再根据具体情况实现细节
	// type interface_name interface{
	//   method_01(参数列表){返回值列表}
	//   ...
	//   ...
	// }
	// 说明：接口里面的所有方法都没有方法体，体现了多态和高内聚低耦合的思想
	// go中的接口不需要显示的实现，只要一个变量，含有接口类型中的所有方法，那么这个变量就实现了接口

	// 接口的注意事项：
	// 1.接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
	// 2.接口中所有的方法都没有方法体，都是没有实现的方法
	// 3.一个自定义类型需要将某个接口的所有方法都实现，这个自定义类型就实现了该接口
	// 4.一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋给这个接口
	// 5.只要是自定义数据类型就可以实现接口，不仅是结构体类型
	// 6.一个自定义类型可以实现多个接口
	// 7.接口中不能含有任何变量
	// 8.一个接口（A）可以继承多个别的接口（B,C），此时如果要实现A接口，就必须将B,C的接口都全部实现
	// 9.interface默认是指针（引用类型），如果没有对interface初始化，默认值是nil
	// 10.空接口interface没有任何方法，所以所有类型都实现了空接口，空接口可以接收任何数据类型 var demo interface{}

}
