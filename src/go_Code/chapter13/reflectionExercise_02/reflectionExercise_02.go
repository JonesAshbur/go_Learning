package main

import (
	"fmt"
	"reflect"
)

// 定义Cal结构体
type Cal struct {
	Num1 int
	Num2 int
}

// GetSub方法
func (c Cal) GetSub(name string) {
	result := c.Num1 - c.Num2
	fmt.Printf("%s完成了减法运行，%d - %d = %d\n", name, c.Num1, c.Num2, result)
}

func main() {
	// 创建Cal结构体实例
	cal := Cal{Num1: 8, Num2: 3}

	// 使用反射遍历Cal结构体的字段信息
	val := reflect.ValueOf(cal)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Printf("字段名: %s, 字段值: %v\n", typ.Field(i).Name, field.Interface())
	}

	// 使用反射调用GetSub方法
	method := reflect.ValueOf(cal).MethodByName("GetSub")
	if method.IsValid() {
		args := []reflect.Value{reflect.ValueOf("tom")}
		method.Call(args)
	} else {
		fmt.Println("方法未找到")
	}
}
