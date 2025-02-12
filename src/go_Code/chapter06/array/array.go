package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 数组可以存放多个同一数据类型，数组是值类型
	// 元素默认值：0
	// 数组在内存空间中是连续分配的，数组名指向的地址就是数组元素的首地址，以此类推加元素所占用的字节数
	// %p在格式化输出中代表地址

	// var arr1 [3]float64
	// for i := 0; i < len(arr1); i++ {
	// 	fmt.Printf("请输入第%d个元素的值", i)
	// 	fmt.Scanln(&arr1[i])
	// }
	// for i := 0; i < len(arr1); i++ {
	// 	fmt.Printf("数组元素%d：[%d]=%v\n", i, i, arr1[i])
	// }

	// 四种数组元素初始化的方式：
	// 1.var arr2 [3]int = [3]int {1,2,3}
	// 2.var arr3 [3]int = {1,2,3}
	// 3.var arr4 = [...]int {1,2,3}
	// 4.var arr5  = [3]int{1:2,0:1,2:3}

	var arr2 [3]int = [3]int{1, 2, 3}
	var arr3 = [3]int{1, 2, 3}
	var arr4 = [...]int{1, 2, 3}
	var arr5 = [3]int{1: 2, 0: 1, 2: 3}
	fmt.Println("arr2=", arr2)
	fmt.Println("arr3=", arr3)
	fmt.Println("arr4=", arr4)
	fmt.Println("arr5=", arr5)

	// for range遍历数组
	// 返回值index是数组的下标，value是下标所对应元素的值，都是for循环的局部变量，不需要下标可以使用下划线忽略
	// index，value名称自定

	str1 := [...]string{"tom", "jerry", "jack"}
	for index, value := range str1 {
		fmt.Printf("index=%d,value=%v\n", index, value)
	}

	// 数组使用的注意事项：
	// 1.数组一旦声明/定义，数组长度不可改变，不能动态变化
	// 2.var arr []int 不是数组，是切片
	// 3.数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用
	// 4.数组创建后如果没有赋值，默认是零值，数值类型默认是0，字符串默认空串""，bool类型默认false
	// 5.数组越界，报错panic
	// 6.数组是值类型，默认值传递，只会进行值拷贝
	// 7.数组长度是数组数据类型的一部分，长度不同，数据类型不同，传递函数参数时候要考虑数组长度

	// 数组反转
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	// 创建一个长度为5的数组
	var arr [5]int
	// 随机生成五个数并放入数组
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) // 生成0到99之间的随机数
	}
	// 打印原始数组
	fmt.Println("原始数组:", arr)
	// 反转数组
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	// 打印反转后的数组
	fmt.Println("反转后的数组:", arr)
}
