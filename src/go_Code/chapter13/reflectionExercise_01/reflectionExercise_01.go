package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectExercise_01(receive interface{}) {
	receiveVal := reflect.ValueOf(receive)
	fmt.Printf("receiveVal_Kind:%v\n", receiveVal.Kind())
	receiveVal.Elem().SetInt(20)
}
func reflectExercise_02(receive interface{}) {
	receiveVal := reflect.ValueOf(receive)
	fmt.Printf("receiveVal_Kind:%v\n", receiveVal.Kind())
	receiveVal.Elem().Field(0).SetString("jack")
}
func main() {

	// 通过反射修改变量的值
	var num int = 10
	reflectExercise_01(&num)
	fmt.Println("num=", num)
	var stu Student = Student{"tom", 18}
	reflectExercise_02(&stu)
	fmt.Println("stu=", stu)
}
