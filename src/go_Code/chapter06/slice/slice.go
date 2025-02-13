package main

import "fmt"

func main() {

	// 切片（引用类型），是数组的一个引用，切片的使用和数组类似
	// 切片的长度可以变化，因此切片是一个动态数组
	// slice：首元素地址，长度，容量
	var arr1 [5]int = [...]int{1, 2, 3, 4, 5}
	// slice声明一个切片，arr表示引用到这个数组，[1:3]表示下标为1到3（不包含3）
	slice1 := arr1[1:3]
	fmt.Println("arr1:", arr1)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice1长度：", len(slice1))
	fmt.Println("slice1容量：:", cap(slice1)) //切片容量可以动态变化

	// 切片的三种使用方式：
	// 1.定义一个切片然后让其引用创建好的数组
	// 2.通过make来创建切片 var 切片名[] type = make([],len,[cap]) 默认值全部为零，cap是可选项，cap>=len
	// 对于切片要先make再使用
	// 通过make方式创建的切片，对应的数组是由make底层维护，对外不可见，即只能通过slice访问
	// 3.定义一个切片，直接就指定具体数组，使用原理类似make

	// 方式1和方式2的区别：方式一是直接引用一个数组，数组事先存在，可见
	// 方式二是通过make创建切片，make也会创建一个数组，	由切片在底层进行维护，不可见
	var slice2 []int32 = make([]int32, 3, 5)
	slice2[0] = 1
	slice2[1] = 2
	fmt.Println("slice2:", slice2)
	fmt.Println("slice2长度：", len(slice2))
	fmt.Println("slice2容量：", cap(slice2))

	// 方式3：
	var slice3 []string = []string{"tom", "jack", "jones"}
	fmt.Println("slice3:", slice3)
	fmt.Println("slice3长度：", len(slice3))
	fmt.Println("slice3容量：", cap(slice3))

	// 切片的遍历
	// for循环，for range
	var arr2 [5]float32 = [...]float32{1.2, 2.3, 3.4, 4.5, 5.6}
	var slice4 = arr2[0:3]
	for i := 0; i < len(slice4); i++ {
		fmt.Printf("slice4[%d]=%v\n", i, slice4[i])
	}

	for index, value := range slice4 {
		fmt.Printf("index:%v,value=%v\n", index, value)
	}

	// 切片使用的注意事项：
	// 1.切片初始化时：var slice = arr[startIndex:endIndex]数组下标为startIndex，取到下标为endIndex（不含endIndex）
	// 2.切片初始化不能越界，范围[0,len(arr)],可以动态增长
	// 3.var slice = arr[:end] 从0开始，0可以省略
	// 4.var slice = arr[start:len(arr)]可以简写成 var slice = arr[start:]
	// 5.var slice = arr[0:len(arr)]可以简写成var slice = arr[:]
	// cap是一个内置函数，用于统计切片容量，即最大可以存储多少个元素
	// 切片定义完成后不能直接使用，本身为空，需要引用一个数组，或者使用make
	// 切片可以继续切片

	// append内置函数可以对切片进行动态追加
	var slice5 []int = []int{100, 200, 300}
	fmt.Println("slice5=", slice5)
	slice5 = append(slice5, 400, 500)
	fmt.Println("slice5=", slice5)
	slice5 = append(slice5, slice5...)
	fmt.Println("slice5=", slice5)
	// append原理：
	// 切片append操作本质上是对数组扩容
	// go底层会创建新的数组newArray
	// 将原来slice包含的元素拷贝到新的数组newArray
	// slice重新引用新的数组newArray
	// newArray在底层进行维护，不可见

	// 切片的拷贝：切片使用内置函数copy完成拷贝,拷贝类型必须都是切片
	// 如果slice容量不够，则根据有效容量进行拷贝
	var arr3 []int = []int{1, 2, 3, 4, 5}
	var slice6 = make([]int, 10)
	fmt.Println("slice6=", slice6)
	copy(slice6, arr3)
	fmt.Println("slice6=", slice6)

	// string和slice
	// string底层是一个byte数组，因此string也可以进行切片处理
	// string是不可改变的，不能通过slice来修改string
	str1 := "hello,golang"
	slice7 := str1[6:]
	fmt.Println("slice7:", slice7)
	// 如果需要修改字符串，先将字符串转换成一个byte切片[]byte或者[]rune，然后进行修改
	str2 := "golang"
	slice8 := []byte(str2)
	slice8[0] = 'G'
	str2 = string(slice8)
	fmt.Println("str2:", str2)
	// 转成[]byte后可以修改数字和英文，但是不能修改中文，因为byte占一个字节，而中文字符占三个字节
	// 解决方法是将带有中文字符的字符串转成[]rune切片，rune按照字符处理，兼容汉字
	str3 := "你好"
	slice9 := []rune(str3)
	slice9[0] = '不'
	str3 = string(slice9)
	fmt.Println("str3:", str3)
}
