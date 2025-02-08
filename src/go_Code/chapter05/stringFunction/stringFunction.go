package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 统计字符串长度，按字节len(str)
	// 中文字符占三个字节
	str1 := "你好golang"
	fmt.Println("str1长度：", len(str1))

	// 字符串的遍历，同时处理含有中文的问题
	str2 := []rune(str1)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("字符：%c\n", str2[i])
	}

	// 字符串转整数：n,err := strconv.Atoi("123")
	res, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果：", res)
	}

	// 整数转字符串
	str3 := strconv.Itoa(123)
	fmt.Printf("str3=%v,str3数据类型：%T\n", str3, str3)

	// 字符串转[]byte:var bytes = []byte("golang")
	var bytes = []byte("golang")
	fmt.Printf("bytes=%v,bytes数据类型：%T\n", bytes, bytes)

	// []byte 转字符串：str := string([]byte{97,97,97,97})
	str4 := string([]byte{97, 98, 99, 100})
	fmt.Println("str4=", str4)

	// 十进制转2，8，16 str := strconv.FormatInt(5,2)
	str5 := strconv.FormatInt(5, 2)
	fmt.Println("str5=", str5)

	// 查找子串是否在指定的字符串中strings.Contains("golang","go")	返回布尔值
	res1 := strings.Contains("golang", "go")
	fmt.Println("res1=", res1)

	// 统计字符串中有几个不重复的指定字串 strings.Count("hello","h")
	res2 := strings.Count("helloGooolang", "o")
	fmt.Println("res2=", res2)

	// 不区分字母的小写的字符串比较（==区分大小写） strings.EqualFold("abc","ABC") 返回布尔值
	res3 := strings.EqualFold("abc", "ABC")
	fmt.Println("res3=", res3)
	fmt.Println("abc" == "ABC")

	// 返回字串在指定字符串中第一次出现的index值 strings.Index("hello golang","go") 找不到返回-1
	// 返回最后一个：strings.LastIndex
	res4 := strings.Index("hello golang", "go")
	fmt.Println("res4=", res4)

	// 将指定的子串替换成另一个子串strings.Replace("golang","go","Go"),数字表示替换个数，-1表示全部替换,原字符串可以是变量，原字符串不变
	res5 := strings.Replace("golang go", "go", "Go", 2)
	fmt.Println("res5=", res5)

	// 按照指定字符作为分割符，将字符串拆分成字符串数组 strings.Split("hello golang","o")
	res6 := strings.Split("hello,golang", ",")
	for i := 0; i < len(res6); i++ {
		fmt.Printf("res6[%v]=%v\n", i, res6[i])
	}
	fmt.Println("res6=", res6)

	// 字符串大小写转换，strings.ToLower("Go"),strings.ToUpper("go")
	res7 := strings.ToLower("GO")
	res8 := strings.ToUpper("go")
	fmt.Println("res7=", res7)
	fmt.Println("res8=", res8)

	// 去掉字符串左右两边空格 strings.TrimSpace(" hello golang ")
	res9 := "  hello golang  "
	res10 := strings.TrimSpace(res9)
	fmt.Printf("res9=%q\n", res9)
	fmt.Printf("res10=%q\n", res10)

	// 将字符串两边指定的字符去掉 strings.Trim("hello golang","hello"),左边TrimLeft,右边TrimRight
	res11 := strings.Trim(" ?hello golang? ", " ?")
	fmt.Printf("res11=%q\n", res11)

	// 判断字符串是否以某个字符串开头strings.HasPrefix("https://www.google.com","http1")
	res12 := strings.HasPrefix("https://www.google.com", "http1")
	fmt.Println("res12=", res12)

	// 判断字符串是否以某个字符串结束strings.HasSuffix("https://www.google.com",".com")
	res13 := strings.HasSuffix("https://www.google.com", ".com")
	fmt.Println("res13=", res13)
}
