package main

// 包前面加下划线 _ 表示忽略此包，但是不删除
import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	// int8,int16,int32,int64,默认int是有符号的，32位系统占4个字节，64位系统占8个字节
	// int32别名为rune，int8别名为byte
	// uint无符号数，32位系统占4个字节，64位系统占8个字节
	// rune有符号数，等价于int32，表示一个unicode码
	// byte用来存储字符，无符号，与uint8等价

	// 整型的注意事项：
	// 1.整型默认是int类型
	// 2.在格式化输出fmt.Printf()中使用 %T 可以输出变量的数据类型
	// 3.保证程序正常运行的情况下，使用占用内存空间小的数据类型

	var num1 = 10
	fmt.Printf("num1的数据类型是：%T \n", num1)
	var num2 int64 = 20
	fmt.Printf("num2的数据类型是：%T ,num2占用的字节数是：%d \n", num2, unsafe.Sizeof(num2))

	// 浮点型的注意事项：
	// 1.分类：单精度浮点型float32(4 bytes)，双精度浮点型float64(8 bytes)
	// 2.浮点数 = 符号位 + 指数位 + 尾数位
	// 3.尾数部分可能丢失，造成精度损失
	// 4.Golang的浮点类型有固定的范围和字段长度，不受os影响
	// 5.浮点型默认声明为float64
	// 6.两种表现形式：十进制表示（必须有小数点），科学计数法e/E都可

	num3 := 1.2
	num4 := .123 //等价于0.123
	num5 := 5.12e2
	fmt.Println("num3=", "num4=", "num5=", num3, num4, num5)

	var num6 = 1.1
	fmt.Printf("num6的数据类型是：%T \n", num6)

	// 字符类型的注意事项：
	// 1.Golang没有专门的字符类型，存储单个字符用byte
	// 2.Golang的字符串由单个字节连接起来（由字节组成）
	// 3.如果保存的字符包含在ASCII码表可用byte保存，如果超出（>255）则使用int
	// 4.字符常量用单引号引起来
	// 5.Go语言字符使用UTF-8编码，英文字母1个字节，汉字3个字节
	// 6.字符的本质是一个整数，直接输出的是该字符的ASCII码值
	// 7.变量赋值数字，按照格式化输出也可以输出该值对应的字符
	// 8.字符类型可以进行运算

	var char1 byte = 'a'
	fmt.Println("char1=", char1)
	fmt.Printf("char1= %c ", char1)
	var char2 int = 11127
	fmt.Printf("char2= %c \n", char2)

	// 布尔类型的注意事项：
	// 1.bool类型的两个值：true,fasle(不可使用0和非0数代替)
	// 3.bool类型占一个字节

	// 字符串类型的注意事项：
	// 1.Golang字符串使用utf-8编码，解决中文乱码
	// 2.字符串一旦赋值不可修改
	// 3.两种表现形式：双引号（会识别转义字符），反引号（以字符串的原生形式输出，包括换行和特殊字符，可实现防攻击，输出源代码）
	// 4.加号可实现字符串的拼接，一行太长时可以使用多行，将 + 留在上一行

	var address string = "中国首都北京"
	fmt.Println(address)

	// 引号
	var str1 = "abc\nabc"
	fmt.Println(str1) //有换行操作，将\n转义

	// 反引号
	str2 := `abc\nabc`
	fmt.Println(str2) //没有将\n转义

	// 拼接字符串
	str3 := "hello " + "golabl "
	str3 += "world"
	fmt.Println(str3)

	// 读行字符串拼接
	str4 := "hello " + "hello " +
		"hello " + "hello " + "hello " +
		"hello " + "hello " + "hello " +
		"hello " + "hello " + "hello " +
		"hello " + "hello " + "hello "
	fmt.Println(str4)

	// 基本数据类型默认值（又叫零值）：
	// 1.整型：0
	// 2.浮点型：0
	// 3.字符串：""
	// 4.布尔类型：false
	var str5 string
	fmt.Println(str5)
	var boolean1 bool
	fmt.Println(boolean1)

	// fmt.Printf("%v") %v是按照 原值 进行输出

	// 基本数据类型的转换：
	// 1.Golang数据类型转换只能显式（强制）转换，不能自动（隐式）转换
	// 2.表达式：DataType(Variable)
	// 3.数据类型转换的表示范围也可以修改：小->大，大->小
	// 4.被转换的是变量的值，变量本身的数据类型不会发生转换
	// 5.大到小转换时，变量的值超出范围时，按照溢出处理，编译不会报错

	var num7 int32 = 100
	var num8 float32 = float32(num7)
	fmt.Printf("num7=%d, num8=%.2f, num8的数据类型是：%T\n", num7, num8, num8)

	var num9 int64 = 29392739729739
	var num10 int8 = int8(num9)
	fmt.Printf("num9=%d,num10=%d\n", num9, num10) //发生溢出

	// 基本数据类型转成string
	// 1.fmt.Sprintf("%参数",表达式)，%q 输出的内容自动加双引号
	// 2.使用strconv包中的函数

	var num11 int8 = 10
	var num12 float32 = 20.0
	var boolean2 bool = true
	var str6 byte = 'a'
	var str7 string

	str7 = fmt.Sprintf("%d", num11)
	fmt.Printf("str7的数据类型：%T,值为：%v\n", str7, str7)

	str7 = fmt.Sprintf("%f", num12)
	fmt.Printf("str7的数据类型：%T,值为：%v\n", str7, str7)

	str7 = fmt.Sprintf("%t", boolean2)
	fmt.Printf("str7的数据类型：%T,值为：%v\n", str7, str7)

	str7 = fmt.Sprintf("%c", str6)
	fmt.Printf("str7的数据类型：%T,值为：%v\n", str7, str7)

	str8 := strconv.FormatInt(int64(num11), 10)
	fmt.Printf("str8的数据类型：%T,值为：%v\n", str8, str8)

	// f 为输出格式，10表示小数位数保留十位，64代表小数是float64
	str8 = strconv.FormatFloat(float64(num12), 'f', 10, 64)
	fmt.Printf("str8的数据类型：%T,值为：%v\n", str8, str8)

	str8 = strconv.FormatBool(boolean2)
	fmt.Printf("str8的数据类型：%T,值为：%v\n", str8, str8)

	// Itoa函数
	var num13 int64 = 100
	var str9 string = strconv.Itoa(int(num13))
	fmt.Printf("str9的数据类型：%T,值为：%v\n", str9, str9)

	// string转基本数据类型
	// 1.使用strconv包里的函数
	// 2.函数会返回两个值第一个是value，第二个是error，下划线 _ 忽略第二个值
	// 3.将string转成其它基本数据类型时要确保能够转成有效的数据，如果非法则直接转成0,布尔类型转为false

	var str_example string = "hello"
	var num_example int64
	num_example, _ = strconv.ParseInt(str_example, 10, 64)
	fmt.Printf("num_example的数据类型：%T,值为：%v\n", num_example, num_example)

	var str10 string = "true"
	var boolean3 bool
	boolean3, _ = strconv.ParseBool(str10)
	fmt.Printf("boolean3的数据类型：%T,值为：%v\n", boolean3, boolean3)

	var str11 = "12345"
	var num14 int64
	num14, _ = strconv.ParseInt(str11, 10, 64)
	fmt.Printf("num14的数据类型：%T,值为：%v\n", num14, num14)

	var str12 = "123.456"
	var num15 float64
	num15, _ = strconv.ParseFloat(str12, 64)
	fmt.Printf("num15的数据类型：%T,值为：%v\n", num15, num15)

	// 因为返回的结果是int64和float64，如果需要int32和float32则使用如下方法
	// var num int32
	// num = int32(variable)

	// 指针类型
	var num16 int8 = 100
	fmt.Println("num16的地址是：", &num16)
	var ptr1 *int8 = &num16
	fmt.Printf("ptr1=%v\n", ptr1)
	fmt.Printf("ptr1=%v\n", *ptr1)
	fmt.Printf("ptr1=%v\n", &ptr1)
	// 值类型都有对应的指针类型，形式为 *数据类型
	// 值类型包括：int，float，bool，string，数组，结构体
	// 引用类型：指针，切片slice，map，管道，interface
	// 区别：值类型变量直接存储值，内存在栈中分配，引用类型变量存储的是一个地址，地址对应的空间才是真实值，内存通常在堆上分配，
	//       当没有任何一个变量来引用这个地址时，地址对应的数据空间变为垃圾，由GC回收
	// uintptr是一个整型，但是它的长度是不确定的，但是在不同的平台上，它的长度是相同的，存放的是一个地址，可以计算
	// go语言支持自定义数据类型
	// type 自定义数据类型名 数据类型 //相当于别名
	// ex：type myint8 int8   //此时myint8可以当作int8使用
	// ex：type mySum func(int,int) int	//mySum等价于 func(int,int) int
	// 即使自定义数据类型和数据类型作用相同，但是数据类型不同
	// myint8 虽然底层是 int8，但在 Go 的类型系统中，它被视为一个全新的独立类型
	// 自定义数据类型可以提高代码可读性
	type myint8 int8
	var num17 myint8 = 100
	fmt.Printf("myint8的数据类型是：%T,num17=%v\n", num17, num17)
	// var num18 int8 = 5
	// var num19 myint8 = num18 // 编译错误,需要显式类型转换：num19 = myint8(num18)
}
