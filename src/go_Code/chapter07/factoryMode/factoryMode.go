package main

import (
	"fmt"

	"github.com/jonesashbur/go_Learning/src/go_Code/chapter07/model"
)

func main() {

	// golang的结构体中没有构造函数，可以使用工厂模式来解决
	// 场景：当某一个包内的结构体变量首字母是小写，但是需要在另一个包中创建这个结构体的实例

	// 创建Student实例(首字母大写)
	stu_01 := model.Student_01{
		Name: "tom",
		Age:  25,
	}
	fmt.Println("stu_01:", stu_01)

	// 创建student实例(首字母小写)
	var stu_02 = model.NewStudent("mary", 18)
	// 整体提取
	fmt.Println("stu_02:", *stu_02)
	// 单独字段提取,*号可有可无，编译器已经优化
	fmt.Println("stu_02_name:", (*stu_02).Name)
	fmt.Println("stu_02_age:", stu_02.GetAge())
}
