package main

import (
	"fmt"
)

func main() {

	// 分支控制：单分支，双分支，多分支
	// fmt.Println("今天学习了吗？(输入true或者false)")
	// var isStudy bool
	// fmt.Scanln(&isStudy)
	// if isStudy {
	// 	fmt.Println("good job")
	// } else {
	// 	fmt.Println("bad job")
	// }

	// ex:golang支持在if中直接定义一个变量
	if num1 := 10; num1 > 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
