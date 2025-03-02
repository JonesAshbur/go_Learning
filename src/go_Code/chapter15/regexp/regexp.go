package main

import (
	"fmt"
	"regexp"
)

func RegexpCase() {
	//	构建正则表达式对象
	reg := regexp.MustCompile(`^[a-z]+\[[0-9]+]$`)
	//	判断给定字符串是否符合正则
	fmt.Println(reg.MatchString("abc[1234]"))
	//	从给定的字符串查找符合条件的字符串
	bytes := reg.FindAll([]byte("efg[4567]"), -1)
	fmt.Printf("符合条件的字符串为：%s\n", bytes[0])
	fmt.Println("符合条件的字符串为：", string(bytes[0]))
}
func main() {
	RegexpCase()
}
