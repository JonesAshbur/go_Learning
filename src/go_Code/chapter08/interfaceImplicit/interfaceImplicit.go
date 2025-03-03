// golang对象实现interface无需关键词，只需要该对象的方法集中包含接口定义的所有方法并且方法签名一致
// 对象实现接口可以借助struct内嵌特性，实现接口的默认实现
// 类型T方法集包含全部receiver T 方法；类型*T方法集包含receive T + *T方法
// 类型T实例value或者pointer都可以调用全部方法，编译器会自动转换
// 类型*T实现接口，只有T类型的指针实现了该接口
// 类型T实现接口，T和*T都实现了该接口
package main

import "fmt"

// 接口
type Animal interface {
	Eat()
	Drink()
	Sleep()
	Run()
}

// animal实例
type animal struct {
}

// animal实现了接口
func (a animal) Eat() {
	fmt.Println("Animal Eat接口默认实现")
}
func (a *animal) Drink() {
	fmt.Println("Animal Drink接口默认实现")
}
func (a animal) Sleep() {
	fmt.Println("Animal Sleep接口默认实现")
}
func (a animal) Run() {
	fmt.Println("Animal Run接口默认实现")
}

func init() {
	animal := animal{}
	animal.Eat() //理论上只有eat方法，但是drink方法也可以使用是因为编译器会自动转换
	animal.Drink()
	animal.Sleep()
	animal.Run()
}

// cat继承了animal的全部方法
type Cat struct {
	animal
}

func NewCat() Animal {
	return &Cat{}
}

func AnimalLife(animal Animal) {
	animal.Eat()
	animal.Drink()
	animal.Sleep()
	animal.Run()
}

// go支持方法重写不支持方法重载
func (cat *Cat) Eat() {
	fmt.Println("重写专属于猫的方法（Eat）")
}
func main() {
	cat := NewCat()
	AnimalLife(cat)
}
