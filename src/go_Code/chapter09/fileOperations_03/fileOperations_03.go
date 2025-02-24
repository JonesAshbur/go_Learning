package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 接收文件函数
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, error := os.Open(srcFileName)
	if error != nil {
		fmt.Println("打开文件错误：", error)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	dstFile, error := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666) // 修改了打开模式
	if error != nil {
		fmt.Println("打开dst文件错误：", error)
		return // 加上 return 以便在发生错误时退出函数
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}
func main() {

	// 文件拷贝
	srcFilePath := "E:/Pictures/photo-1544954773-3793881f46c6.jfif"
	dstFilePath := "E:/Pictures/demo.jpg"
	_, error := CopyFile(dstFilePath, srcFilePath)
	if error == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println("拷贝失败：", error)
	}
}
