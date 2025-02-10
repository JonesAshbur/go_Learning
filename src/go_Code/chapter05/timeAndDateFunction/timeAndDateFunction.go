package main

import (
	"fmt"
	"strconv"
	"time"
)

func splicing() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "Golang" + strconv.Itoa(i)
	}
}

func main() {

	// 日历计算采取公历
	nowTime := time.Now()
	fmt.Printf("现在的时间是：%v\n 数据类型是：%T\n", nowTime, nowTime)
	fmt.Printf("年:%v\n", nowTime.Year())
	fmt.Printf("月:%v\n", int(nowTime.Month()))
	fmt.Printf("日:%v\n", nowTime.Day())
	fmt.Printf("时:%v\n", nowTime.Hour())
	fmt.Printf("分:%v\n", nowTime.Minute())
	fmt.Printf("秒:%v\n", nowTime.Second())

	// 格式化日期时间方式一：
	fmt.Printf("当前年月日：%d-%d-%d-%d-%d-%d\n", nowTime.Year(), nowTime.Month(), nowTime.Day(),
		nowTime.Hour(), nowTime.Minute(), nowTime.Second())

	dataString := fmt.Sprintf("当前年月日：%d-%d-%d-%d-%d-%d\n", nowTime.Year(), nowTime.Month(), nowTime.Day(),
		nowTime.Hour(), nowTime.Minute(), nowTime.Second())
	fmt.Printf("dataString=%v数据类型：%T\n", dataString, dataString)

	// 格式化日期时间方式二：数字固定不可修改，各个数字可以自由组合
	fmt.Printf("%s", nowTime.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf("%s", nowTime.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf("%s", nowTime.Format("15:04:05"))
	fmt.Println()

	// 时间的常量
	// 作用：在程序中获取指定时间单位的时间
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )

	// 每隔一秒打印一个数字，打印到100退出
	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠
		time.Sleep(time.Microsecond * 100)
		if i == 100 {
			break
		}
	}

	// 获取当前unix时间戳和unixnao时间戳（可以获取随机数字）
	fmt.Printf("unix时间戳：%v\nunixnao时间戳：%v\n", nowTime.Unix(), nowTime.UnixNano())

	start := time.Now().Unix()
	splicing()
	end := time.Now().Unix()
	fmt.Printf("执行splicing花费秒数：%v\n", end-start)
}
