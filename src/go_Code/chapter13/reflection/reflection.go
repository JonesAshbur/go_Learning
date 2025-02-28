package main

import (
	"fmt"
	"reflect"
)

func reflectTest_01(receive interface{}) {
	// 先获取reflect.Type
	reflectType := reflect.TypeOf(receive)
	fmt.Println("reflectType=", reflectType)
	fmt.Printf("reflectType=%T\n", reflectType)
	// 获取reflect.Value
	reflectValue := reflect.ValueOf(receive)
	fmt.Println("reflectValue=", reflectValue)
	fmt.Printf("reflectValue=%T\n", reflectValue)
	var demo int64 = 100
	demo += reflectValue.Int()
	fmt.Println("demo + reflectValue = ", demo)
	// 将reflectValue转回interface{}
	res := reflectValue.Interface()
	fmt.Printf("res=%v\n", res)
	fmt.Printf("res=%T\n", res)
	// 将interface转成需要的数据类型
	num := res.(int)
	fmt.Printf("num=%v\n", num)
	fmt.Printf("num=%T\n", num)
}

type Student struct {
	Name string
	Age  int
}

func reflectTest_02(receive interface{}) {
	// 先获取reflect.Type
	reflectType := reflect.TypeOf(receive)
	fmt.Println("reflectType=", reflectType)
	fmt.Printf("reflectType=%T\n", reflectType)
	// 获取reflect.Value
	reflectValue := reflect.ValueOf(receive)
	fmt.Println("reflectValue=", reflectValue)
	fmt.Printf("reflectValue=%T\n", reflectValue)
	// 获取kind
	typekind := reflectType.Kind()
	fmt.Println("typekind=", typekind)
	valuekind := reflectValue.Kind()
	fmt.Println("valuekind=", valuekind)
	// 将reflectValue转回interface{}
	res := reflectValue.Interface()
	fmt.Printf("res=%v\n", res)
	fmt.Printf("res=%T\n", res)
	// 将interface转成需要的数据类型
	stu := res.(Student)
	fmt.Printf("stu_Name=%v\n", stu.Name)
	fmt.Printf("stu_Age=%v\n", stu.Age)

}
func main() {

	// 反射可以在程序运行时动态的获取变量的各种信息，比如变量的类型和类别
	// 如果是结构体变量，还可以获取到结构体本身的各种信息（字段，方法）
	// 通过反射可以修改变量的值，可以调用关联的方法
	// 使用反射，import "reflect"
	// 变量，interface{}，reflect.value可以相互转换

	var num_01 int = 100
	reflectTest_01(num_01)

	stu := Student{
		Name: "Tom",
		Age:  18,
	}
	reflectTest_02(stu)

	// 反射的注意事项和细节
	// 1. reflect.Value.Kind()获取变量的类别，返回的是一个常量
	// 2. Type是类型，Kind是类别，Type和Kind可能相同，也可能不同
	// 3. 通过反射可以让变量在interface{}和reflect.Value之间相互转换
	// 4. 使用反射的方式来获取变量的值（并返回对应的类型）要求数据类型匹配

	//方法的排序默认是按照函数名的ASCII码进行排序

}
