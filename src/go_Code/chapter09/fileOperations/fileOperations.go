package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// 文件操作
	// 输入流：数据从文件到内存的路径
	// 输出流：数据从内存到文件的路径
	// os.File封装所有文件相关的操作，File是一个结构体

	// 打开文件
	// file, error := os.Open("E:/go_Learning/fileDemo")
	// if error != nil {
	// 	fmt.Println("打开文件失败，原因是：", error)
	// } else {
	// 	fmt.Println("文件内容是：", file)
	// }
	// // 关闭文件
	// err := file.Close()
	// if err != nil {
	// 	fmt.Println("关闭文件错误", err)
	// } else {
	// 	fmt.Println("关闭成功")
	// }

	// 读取文件内容并且显示在终端（带缓冲区的形式）
	file, error := os.Open("E:/go_Learning/fileDemo/test.txt")
	if error != nil {
		fmt.Println("打开文件失败，原因是：", error)
	}
	// 关闭文件
	defer file.Close() //及时关闭文件句柄，避免内存泄漏
	// 创建一个带缓冲的*reader，默认缓冲区为4096
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到换行就结束
		if err == io.EOF {                  //io.EOF代表读到文件末尾
			break
		}
		fmt.Print(str)
	}

	fmt.Println("文件读取结束...")

	// 读取文件内容并且显示在终端（使用ioutil一次性将整个文件读取到内存中），这种方式适用于文件不大的情况
	filepath_01 := "E:/go_Learning/fileDemo/test.txt"
	content, err := os.ReadFile(filepath_01) // 使用 os.ReadFile
	if err != nil {
		fmt.Println("读取文件失败:", err)
	}
	fmt.Println("文件内容:", string(content))

	// 写文件操作
	// func OpenFile(name string,flag int,perm FileMode) (file *File,err error)
	// oe.OpenFile是一个更一般性的文件打开函数，会使用指定的选项，指定的模式打开指定名称的文件，成功返回文件对象可用于I/O，出错类型是*PathError
	// name是文件名，flag是文件打开模式（可以通过|进行组合），perm是权限控制（Linux）
	// const (
	// 	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	// 	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	// 	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	// 	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// 	// The remaining values may be or'ed in to control behavior.
	// 	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	// 	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	// 	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	// 	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	// 	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	// )

	// 创建一个新文件
	filepath_02 := "E:/go_Learning/fileDemo/demo_01.txt"
	file_01, error_01 := os.OpenFile(filepath_02, os.O_WRONLY|os.O_CREATE, 0666)
	if error_01 != nil {
		fmt.Println("打开文件失败，原因是：", error_01)
		return
	}
	defer file.Close()
	// 写入五句hello,golang
	str_01 := "hello,golang\n"
	// 使用带缓冲的writer
	writer := bufio.NewWriter(file_01)
	for i := 0; i < 5; i++ {
		writer.WriteString(str_01)
	}
	// 因为writer带缓存，因此在调用writeString时，数据先写入到缓存，所以需要调用Flush方法，将缓存数据真正写入到文件中
	writer.Flush()
	content_01, err_01 := os.ReadFile(filepath_02) // 使用 os.ReadFile
	if err != nil {
		fmt.Println("读取文件失败:", err_01)
	}
	fmt.Println("文件内容:")
	fmt.Println(string(content_01))
}
