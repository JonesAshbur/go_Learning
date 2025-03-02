package main

import (
	"fmt"
	"math"
)

func MathCase() {
	fmt.Println("2的10次方：", math.Pow(2, 10))
	fmt.Println("2为底1024的对数：", math.Log2(1024))
	fmt.Println("两数取最大值：", math.Max(1, 2))
	fmt.Println("向上取整：", math.Ceil(1.49))
	fmt.Println("向下取整：", math.Floor(1.89))
	fmt.Println("90度角的正弦值：", math.Sin(math.Pi/2))
	fmt.Println("1的反正弦值：", math.Asin(1))
}
func main() {
	MathCase()
}
