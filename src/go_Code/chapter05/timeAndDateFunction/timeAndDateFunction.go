package main

import (
	"fmt"
	"time"
)

func main() {

	// 日历计算采取公历
	nowTime := time.Now()
	fmt.Printf("现在的时间是：%v\n 数据类型是：%T\n", nowTime, nowTime)

}
