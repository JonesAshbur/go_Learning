package main

import (
	"fmt"
	"os"
)

func PathExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		fmt.Println("文件存在")
		return true, nil
	}
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
		return false, err
	}
	return false, err
}
func main() {

	// 文件操作案例
	// 将demo_04文件中的内容写入到demo_05文件中
	filePath_01 := "E:/golang_Learning/fileDemo/demo_04.txt"
	filePath_02 := "E:/golang_Learning/fileDemo/demo_05.txt"
	content, error_01 := os.ReadFile(filePath_01) // 使用 os.ReadFile
	if error_01 != nil {
		fmt.Println("读取文件失败:", error_01)
	}
	fmt.Println("文件内容:", string(content))
	error_02 := os.WriteFile(filePath_02, content, 0666)
	if error_02 != nil {
		fmt.Println("写入文件失败:", error_02)
	}

	// 判断文件是否存在
	// os.Stat()返回错误为nil说明文件或文件夹存在
	// 返回的错误类型使用os.IsNotExist()判断为true，说明不存在
	// 如果返回的错误类型为其他类型，则不确定是否存在
	filePath_03 := "E:/golang_Learning/fileDemo/demo_09.txt"
	PathExist(filePath_03)
}
