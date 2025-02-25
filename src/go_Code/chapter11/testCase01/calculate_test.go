package main

import (
	"testing"
)

// 测试用例，测试addUpper函数的正确性
func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		// fmt.Printf("执行错误，期望值：%v，实际值：%v", 55, res)
		t.Fatalf("执行错误，期望值：%v，实际值：%v", 55, res)
	}
	t.Logf("addUpper()执行正确")
}
